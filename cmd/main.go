package main

import (
	"fmt"
	"ghub-differences/fileops"
	"ghub-differences/githubFuncs"
	"ghub-differences/merge"
	"ghub-differences/output"
	"ghub-differences/structs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v45/github"
	"github.com/r3labs/diff/v3"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v3"
)

func main() {
	fileops.CreateEmptyYaml()
	RET_VALUE := 0
	var changelog diff.Changelog

	//Get the env variables
	ACCESS_TOKEN := kingpin.Flag("ACCESS_TOKEN", "Access Token").Envar("ACCESS_TOKEN").Required().String()
	BASE_URL := kingpin.Flag("BASE_URL", "Base URL").Envar("BASE_URL").Required().String()
	REQ_PATH := kingpin.Flag("REQ_PATH", "Req Path").Envar("REQ_PATH").Required().String()
	OVERRIDE_PATH := kingpin.Flag("OVERRIDE_PATH", "Override Path").Envar("OVERRIDE_PATH").Required().String()
	kingpin.Parse()

	context, client := githubFuncs.Authorize(ACCESS_TOKEN, BASE_URL) //Authorize using access token

	base, err := os.ReadFile(fmt.Sprintf("%v", *REQ_PATH)) //read the requirements yaml file

	if err != nil {
		fmt.Printf("\n%v\n", err)
	}

	override, err := os.ReadFile(fmt.Sprintf("%v", *OVERRIDE_PATH)) //read the override yaml
	if err != nil {
		fmt.Printf("\n%v\n", err)
	}

	var RequirementsList, OverrideList structs.YamlStruct

	err = yaml.Unmarshal(base, &RequirementsList) //Unmarshal the requirements list into the struct
	if err != nil {
		log.Fatal("error while unmarshaling reqlist" + err.Error())
	}

	err = yaml.Unmarshal(override, &OverrideList) //Unmarshal the override list into the struct
	if err != nil {
		log.Fatal("error while unmarshaling reqlist" + err.Error())
	}

	for _, currOrg := range RequirementsList.Organizations {
		//Print the organization name
		fmt.Printf("Organization Name: %s\n", currOrg)

		var Repos []*github.Repository

		//Retry logic
		for i := 0; i < 5; i++ {
			//Fetch the repository list of the authorized user according to the organization name
			Repos, _, err = client.Repositories.ListByOrg(context, currOrg, nil)
			if err == nil {
				break
			}
			//Sleep for 1 sec
			time.Sleep(1 * time.Second)
		}

		OverridedRepoList := make(map[string]bool)
		//Get overrided repository list
		for _, currRepo := range Repos {
			for _, v := range *OverrideList.Override {
				if v.RepoName == *currRepo.Name {
					OverridedRepoList[*currRepo.Name] = true
				}
			}
		}

		StandardBranchList := make(map[string]bool) //Extract the branches that are defined in requirements list.
		var getTeamData *github.Repository
		RequiredTeams := make(map[string]bool) //a map for extracting the teams on requirements list
		for _, orgStandards := range *RequirementsList.OrganizationStandards {

			//Find the current organization standards from the requirements list
			if orgStandards.Organization == currOrg {

				//Get branch list from the requirements list
				for _, stdBranch := range orgStandards.Branches {
					StandardBranchList[stdBranch.BranchName] = true
				}

				//Extract the standard Repository rules and write into a file.
				if orgStandards.Repository != nil {
					StandardRepoFile, _ := yaml.Marshal(orgStandards.Repository)
					if err := os.WriteFile("/reqlist-yamls/Repository.yaml", StandardRepoFile, 0777); err != nil {
						panic(err)
					}
				}

				//Extract the standard Branch Rules
				for _, currentBranch := range orgStandards.Branches {
					if orgStandards.Branches != nil {
						StandardBranchFile, _ := yaml.Marshal(currentBranch)
						if err := os.WriteFile("/reqlist-yamls/"+currentBranch.BranchName+".yaml", StandardBranchFile, 0777); err != nil {
							panic(err)
						}
					}
				}

				//Extract the standard Team Permissions
				if orgStandards.Team != nil {
					for _, currentTeam := range orgStandards.Team {
						RequiredTeams[currentTeam.Slug] = true
						StandardTeamFile, _ := yaml.Marshal(currentTeam)
						if err := os.WriteFile("/reqlist-yamls/Team-"+currentTeam.Slug+".yaml", StandardTeamFile, 0777); err != nil {
							panic(err)
						}
					}
				}
			}
		}
		RepoOverrideStatus := make(map[string]bool)
		var getRepoData *github.Repository
		var BranchList []*github.Branch

		//Iterate over repositories
		for _, currRepo := range Repos {

			fmt.Printf("%v\nFull Name: %v\nName: %v\nURL: %v\n", strings.Repeat("-", 150), *currRepo.FullName, *currRepo.Name, *currRepo.URL)

			if OverridedRepoList[*currRepo.Name] { //if the current repo is overrided
				for _, value := range *OverrideList.Override { //traverse through the strudy
					if value.RepoName == *currRepo.Name { //Find the repo name and corresponding override values

						//Extract the Overrided Repository Rules
						if value.Repository != nil {
							RepoOverrideStatus["Repository"] = true
							RepoFile, _ := yaml.Marshal(value.Repository)
							if err := os.WriteFile("/override-yamls/Repository.yaml", RepoFile, 0777); err != nil {
								panic(err)
							}
						}

						//Extract the Overrided Branch Rules
						if value.Branches != nil {
							for _, overridedBranch := range value.Branches {
								RepoOverrideStatus[overridedBranch.BranchName] = true
								OverridedBranchFile, _ := yaml.Marshal(overridedBranch)
								if err := os.WriteFile("/override-yamls/"+overridedBranch.BranchName+".yaml", OverridedBranchFile, 0777); err != nil {
									panic(err)
								}
							}
						}
						//Extract the Team Permissions
						if value.Team != nil {
							for _, overridedTeam := range value.Team {
								RepoOverrideStatus[overridedTeam.Slug] = true
								TeamFile, _ := yaml.Marshal(overridedTeam)
								if err := os.WriteFile("/override-yamls/Team-"+overridedTeam.Slug+".yaml", TeamFile, 0777); err != nil {
									panic(err)
								}
							}
						}

					}
				}

			}

			getRepoTeams, _, err := client.Repositories.ListTeams(context, currOrg, *currRepo.Name, nil) //get the current repositories teams
			if err != nil {
				fmt.Println(err)
			}

			PresentTeams := make(map[string]bool) //a list for ddetermining the teams that are present in both reqlist and github repo

			//copy the keys and set it all to false
			for k := range RequiredTeams {
				PresentTeams[k] = false
			}

			//Determine the present teams
			for i := range RequiredTeams {
				for _, j := range getRepoTeams {
					if i == j.GetSlug() {
						PresentTeams[i] = true
					}
				}
			}

			TeamAPIData := make(map[string]*github.Repository)

			//Get the team permissions
			for slug, bool_value := range PresentTeams {
				if bool_value {
					//Check the team permissions
					getTeamData, _, err = client.Teams.IsTeamRepoBySlug(context, currOrg, slug, currOrg, *currRepo.Name)
					TeamAPIData[slug] = getTeamData
					if err != nil {
						fmt.Println(err)
					}
				}
			}

			//get the branches of repo
			BranchList, _, err = client.Repositories.ListBranches(context, currOrg, *currRepo.Name, nil)
			if err != nil {
				panic("problem in getting branch list")
			}
			//Get the delete_branch_on_merge data of the repositories. (This value is only returned with a Repositories.Get call, so we have to make another API call.)
			getRepoData, _, err = client.Repositories.Get(context, *currRepo.Owner.Login, *currRepo.Name)
			if err != nil {
				panic("problem in getting repository rule")
			}

			base := map[string]interface{}{}
			override := map[string]interface{}{}

			//																Get repository differences
			// read the first yaml file
			StandardRepoRules, err := os.ReadFile("/reqlist-yamls/Repository.yaml")
			if err != nil {
				log.Fatal("error while reading repo yaml" + err.Error())
			}

			if err := yaml.Unmarshal(StandardRepoRules, &base); err != nil {
				fmt.Printf("\nError while unmarshaling standarreporules: %v\n", err)
			}

			// read the override yaml file, create an empty one if it doesn't exist
			OverridedRepoRules, err := os.ReadFile("/override-yamls/Repository.yaml")
			if err != nil {
				err = os.WriteFile("/override-yamls/Repository.yaml", []byte{}, 0777)
				if err != nil {
					log.Fatal("error creating repo yaml" + err.Error())
				}
				OverridedRepoRules, err = os.ReadFile("/override-yamls/Repository.yaml")
				if err != nil {
					fmt.Printf("\nError while reading override repository: %v\n", err)
				}

			}
			//unmarshal the repo rules
			if err := yaml.Unmarshal(OverridedRepoRules, &override); err != nil {
				panic("cannot unmarshal overrided repo rules")
			}

			// merge both yaml data recursively
			base = merge.MergeMaps(base, override)

			//marshal the merged file
			Output, _ := yaml.Marshal(base)
			if err := os.WriteFile("/standard-yamls/MergedRepoRules.yaml", Output, 0777); err != nil {
				panic(err)
			}

			var MergedRepoRules *structs.Repository

			base = map[string]interface{}{}
			override = map[string]interface{}{}

			ReadYamlFile, _ := os.ReadFile("/standard-yamls/MergedRepoRules.yaml")
			if err := yaml.Unmarshal(ReadYamlFile, &MergedRepoRules); err != nil {
				panic("cannot unmarshal mergedreporules")
			}

			//																Get Team differences
			for slug, bool_value := range PresentTeams {
				base = map[string]interface{}{}
				override = map[string]interface{}{}

				//if team data exist,
				if bool_value {
					for apislug, data := range TeamAPIData {
						if slug == apislug {
							fmt.Printf("\nTeam Name: %v\n", apislug)

							//read the standard perms of current team
							StandardTeamRules, _ := os.ReadFile("/reqlist-yamls/Team-" + slug + ".yaml")
							if err := yaml.Unmarshal(StandardTeamRules, &base); err != nil {
								fmt.Printf("\ncannot unmarshal standard team rules: %v\n", err)
							}

							// read the override yaml file, create an empty one if it doesn't exist
							OverridedTeamRules, err := os.ReadFile("/override-yamls/Team-" + slug + ".yaml")
							if err != nil {
								err = os.WriteFile("/override-yamls/Team-"+slug+".yaml", []byte{}, 0777)
								if err != nil {
									log.Fatal("error while creating emptylist" + err.Error())
								}
								OverridedTeamRules, _ = os.ReadFile("/override-yamls/Team-" + slug + ".yaml")
							}
							if err := yaml.Unmarshal(OverridedTeamRules, &override); err != nil {
								fmt.Printf("\ncannot unmarshal overrided team rules: %v\n", err)

							}

							// merge both yaml data recursively
							base = merge.MergeMaps(base, override)

							TeamOutput, _ := yaml.Marshal(base)
							if err := os.WriteFile("/standard-yamls/MergedTeam-"+slug+".yaml", TeamOutput, 0777); err != nil {
								panic(err)
							}

							var MergedTeamRules structs.TeamClassifier

							ReadFile, err := os.ReadFile("/standard-yamls/MergedTeam-" + slug + ".yaml")
							if err != nil {
								fmt.Println(err)
							}
							if err := yaml.Unmarshal(ReadFile, &MergedTeamRules); err != nil {
								fmt.Printf("\ncannot unmarshal merged team rules: %v\n", err)
							}

							//convert the permissions struct into a map to make the diff operation
							MergedPerms := MergedTeamRules.Permissions

							ConversionInput, _ := yaml.Marshal(MergedPerms)
							if err := os.WriteFile("/standard-yamls/convert-"+slug+".yaml", ConversionInput, 0777); err != nil {
								panic(err)
							}

							ConversionOutput, err := os.ReadFile("/standard-yamls/convert-" + slug + ".yaml")
							if err != nil {
								fmt.Println(err)
							}

							var ConvertedTeamRules map[string]bool
							if err := yaml.Unmarshal(ConversionOutput, &ConvertedTeamRules); err != nil {
								fmt.Printf("\ncannot unmarshal merged team rules: %v\n", err)
							}

							GitHubPerms := data.GetPermissions() //extract the permissions as a map from the github response

							changelog, err := diff.Diff(ConvertedTeamRules, GitHubPerms) //create a changelog for differences
							if err != nil {
								fmt.Printf("\nerror: %v\n", err)
							}
							fmt.Printf("\n\tTeam Differences\n")

							RET_VALUE = output.PrintDiffChangeLog(changelog, RET_VALUE) //print the changelog

						}
					}
				} else {
					fmt.Printf("\nTeam '%v' is not present in this repo.\n", slug)
				}
			}

			fmt.Printf("\n\n")

			//Branch diff
			base = map[string]interface{}{}
			override = map[string]interface{}{}

			fmt.Printf("\n\tBranch Protection Differences\n")
			for _, currentBranch := range BranchList {
				fmt.Printf("\nBranch Name: %v\n", currentBranch.GetName())

				if StandardBranchList[currentBranch.GetName()] { //if true

					// read the first yaml file
					standardBranchRules, _ := os.ReadFile("/reqlist-yamls/" + currentBranch.GetName() + ".yaml")
					if err := yaml.Unmarshal(standardBranchRules, &base); err != nil {
						panic(err)
					}

					overridedBranchRules, err := os.ReadFile("/override-yamls/" + currentBranch.GetName() + ".yaml")
					if err != nil {
						err = os.WriteFile("/override-yamls/"+currentBranch.GetName()+".yaml", []byte{}, 0777)
						if err != nil {
							log.Fatal("error while creating branch data yaml" + err.Error())
						}
						overridedBranchRules, _ = os.ReadFile("/override-yamls/" + currentBranch.GetName() + ".yaml")
					}
					if err := yaml.Unmarshal(overridedBranchRules, &override); err != nil {
						panic("cannot unmarshal overrided branch rules")
					}

					// merge both yaml data recursively
					base = merge.MergeMaps(base, override)

					OUTPUT, _ := yaml.Marshal(base)
					if err := os.WriteFile("/standard-yamls/Merged"+currentBranch.GetName()+".yaml", OUTPUT, 0777); err != nil {
						panic(err)
					}

					var MergedBranchRules *structs.BranchesClassifier

					readYamlFile, _ := os.ReadFile("/standard-yamls/Merged" + currentBranch.GetName() + ".yaml")
					if err := yaml.Unmarshal(readYamlFile, &MergedBranchRules); err != nil {
						panic("cannot unmarshal merged ranch rules")
					}

					getBranchProtectionData, _, protectionErr := client.Repositories.GetBranchProtection(context, *currRepo.Owner.Login, *currRepo.Name, currentBranch.GetName())

					convert, _ := yaml.Marshal(getBranchProtectionData)
					if err := os.WriteFile("/standard-yamls/convert"+currentBranch.GetName()+".yaml", convert, 0777); err != nil {
						panic(err)
					}
					var githubRules *structs.Convert
					readconvert, _ := os.ReadFile("/standard-yamls/convert" + currentBranch.GetName() + ".yaml")
					if err := yaml.Unmarshal(readconvert, &githubRules); err != nil {
						panic("cannot unmarshal merged ranch rules")
					}
					var emptyProtectionList *structs.EmptyProtection
					//Check the error message
					if protectionErr != nil && protectionErr.Error() == "branch is not protected" {
						readYamlFile, _ := os.ReadFile("/empty-yaml/empty.yaml")
						if err := yaml.Unmarshal(readYamlFile, &emptyProtectionList); err != nil {
							panic("cannot unmarshal merged ranch rules")
						}
						fmt.Printf("\n***Branch '%v' is not protected***\n", currentBranch.GetName())
						empty, err := diff.Diff(MergedBranchRules.Protection, emptyProtectionList)
						if err != nil {
							fmt.Printf("\nError not protected branch diff: %v\n", err)
						}
						for _, val := range empty {
							RET_VALUE++
							fmt.Printf("\nDifference: %v\n", val.Path)
						}

					} else {
						changelog, err = diff.Diff(MergedBranchRules.Protection, githubRules)
						if err != nil {
							fmt.Printf("\nError while diffing mergedbrancrules w/ github rules: %v\n", err)
						}
						RET_VALUE = output.PrintDiffChangeLog(changelog, RET_VALUE)

					}

				} else {
					fmt.Println("This branch is not defined in the requirements list.")
				}

			}
			// Repo diff
			changelog, err = diff.Diff(MergedRepoRules, getRepoData)
			if err != nil {
				fmt.Printf("\nError: %v\n", err)
			}
			fmt.Printf("\n\tRepository Differences\n")
			RET_VALUE = output.PrintDiffChangeLog(changelog, RET_VALUE)
			fmt.Printf("\n\n")

			//CODEOWNERS check
			_, directoryContent, _, err := client.Repositories.GetContents(context, currOrg, *currRepo.Name, "", nil)
			if err != nil {
				fmt.Println(err)
			}

			isCodeownersExists := false
			for _, i := range directoryContent {
				if *i.Name == "CODEOWNERS" {
					isCodeownersExists = true
					break
				}
			}

			if isCodeownersExists {
				fmt.Printf("\nCODEOWNERS exists\n")
			} else {
				fmt.Println("CODEOWNERS does not exist")
				RET_VALUE++
			}

			fileops.DeleteTempFiles("/override-yamls")
			fileops.DeleteTempFiles("/standard-yamls")
		}

		//Delete the temp dir content
		fileops.DeleteTempFiles("/reqlist-yamls")
	}
	fmt.Printf("\nDifference Amount: %v\n", RET_VALUE)
	os.Exit(RET_VALUE)
}
