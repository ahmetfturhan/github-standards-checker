# GitHub Repository Rules Differences

## Features
This tool fetches the repository settings from GitHub API and compares it with the predefined rules. Then, states the differences if there are any.

- Differences in branch protection or multiple branch protection rules.

- Differences in repository rules. (i.e. Auto delete head branches on merge)

- Checking if the specified team (or multiple) is present with the given permissions, and stating any differences.

- Checking if the CODEOWNERS file exists.

# Running w/Kubernetes
You can initialize ConfigMaps, Secrets, Service Accounts & create a CronJob using this command;
```
kubectl apply -k base
```

# Running w/Docker
You can run as a Docker Container using this command,
```
docker build --tag github-checker .  

docker run -e ACCESS_TOKEN={ACCESS_TOKEN} -e BASE_URL={BASE_URL} -e REQ_PATH={REQ_PATH} -e OVERRIDE_PATH={OVERRIDE_PATH} github-checker
```

# Requirements

requirements-list.yaml file under /config directory. Structure of the YAML file is as follows,
```
Organizations: //All repositories of the specified organization will be checked
    - org_name
    - org_name2

OrganizationStandards:
    Organization: an-organization
        Repository:
            Repository Rules Here

        Branches:
            - branchName: main
              Protection:
                Branch Protection Rules Here

        Team:
            - Slug: test-team
              Permissions:
                  Permission Rules here

    Organization:
        Same structure above

```

override.yaml file is under /config directory. Structure of the YAML file is as follows,

```
Override:
  - Organization: an-organization
    repoName: test-repo-settings
    Repository:
      delete_branch_on_merge: true

    Branches:
      - branchName: main
        Protection:
          allow_force_pushes:
            enabled: false

    - Organization: org_name
      repo_name: another_repo
      *** Same structure above ***

```

## Example requirements-list.yaml

<details>
  <summary>Click to see an example requirementsList.yaml file</summary>
  
  ### Example:
  ```yaml
    Organizations:
      - an-organization

    OrganizationStandards:
      - Organization: an-organization
      
        Repository:
          delete_branch_on_merge: false
          allow_rebase_merge: false

        Team:
          - Slug: test-team
            Permissions:
              admin: false
              maintain: true
              pull: true
              push: false
              triage: true
      
          - Slug: another-team
            Permissions:
              admin: false
              maintain: true
              pull: true
              push: false
              triage: true
                  

        Branches:
          - branchName: main
            Protection:
              allow_deletions:
                enabled: false
              enforce_admins:
                url: asd123
                enabled: true
              allow_force_pushes:
                enabled: true

          - branchName: deneme
            Protection:
              allow_force_pushes:
                enabled: false
              

  ```
</details>


## Example requirements-list.yaml

<details>
  <summary>Click to see an example override.yaml file</summary>

    ```
    Override:
      - Organization: an-organization
        repoName: test-repo-settings
        Repository:
          delete_branch_on_merge: true

        Branches:
          - branchName: main
            Protection:
              enforce_admins:
                url: isthisworks
              allow_force_pushes:
                enabled: false
              allow_deletions:
                enabled: true

  
    ```
  </details>

## Environmental Variables

GitHub Personal Access Token, Base URL, Requirements YAML path and Override YAML path needs to be set as environment variables.
```
    ACCESS_TOKEN = ghp_8zQ...
    BASE_URL = https://github.{YOUR_ENTERPRISE}.com/api/v3
    REQ_PATH = /etc/config/requirements/requirements.yaml
    OVERRIDE_PATH = /etc/config/override/override.yaml
```

# Rules
There are 3 rulesets;
  - Repository
  - Protection
  - Team

### Repository
These are [Repository Rules](https://pkg.go.dev/github.com/google/go-github/v45/github#Repository)  
For example if you want to check the DeleteBranchOnMerge, you need to add the json tag of the corresponding rule to the checklist.
```yaml
Repository:
  delete_branch_on_merge: true
```

### Protection
These are Branch [Protection Rules](https://pkg.go.dev/github.com/google/go-github/v45/github#Protection)  
Example,
```
Protection:
  allow_deletions:
    enabled: true
  allow_force_pushes:
    enabled: false
  branch_restrictions:
    users:
      - login: dicaprio
      - login: hardy
      - login: mcqueen
    teams:
      - slug: team-of-another-organization
    apps:
      - slug: app-name
```

### Team
These are [Permissions of Teams](https://pkg.go.dev/github.com/google/go-github/v45/github#Team)  
Example,
```
Team:
      TeamSlug: test-team
      Permissions:
        admin: true
        maintain: true
        pull: true
        push: true
        triage: true
```
## Examples


### You want to override a rule for a particular repository,

```yaml
#Requirements YAML
Organizations:
    - another-organization

OrganizationStandards:
  - Organization: another-organization
    Repository:
        delete_branch_on_merge: true

#Override YAML
Override:
  - Organization: another-organization
    repo_name: repo1
    Repository:
      delete_branch_on_merge: false

```

### You want to check the rules for two organizations
```yaml
Organizations:
  - sample-org
  - another-org

OrganizationStandards:
 - Organization: sample-org
   Repository:
      delete_branch_on_merge: true

 - Organization: another-org
   Repository:
      delete_branch_on_merge: false
```

