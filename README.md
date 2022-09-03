# GitHub Repository Rules Differences

## Features
This tool fetches the repository settings from GitHub API and compares it with the predefined rules. Then, states the differences if there are any.

- Differences in branch protection or multiple branch protection rules.

- Differences in repository rules. (i.e. Auto delete head branches on merge)

- Checking if the specified team (or multiple) is present with the given permissions, and stating any differences.

- Checking if the CODEOWNERS file exists.

# Requirements

requirements-list.yaml file under /config directory. Structure of the YAML file is as follows,
```
Organizations: //All repositories of the specified organization will be checked
    - org_name
    - org_name2

OrganizationStandards:
    Organization: cloud-interns
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
  - Organization: cloud-interns
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
      - cloud-interns

    OrganizationStandards:
      - Organization: cloud-interns
      
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
      - Organization: cloud-interns
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
    ACCESS_TOKEN = ghp_6zB...
    BASE_URL = https://github.pkgms.com/api/v3
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
      - login: krisztian
      - login: serhat
      - login: ilyas
    teams:
      - slug: team-of-cloud-team
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
    - cloud-team

OrganizationStandards:
  - Organization: cloud-team
    Repository:
        delete_branch_on_merge: true

#Override YAML
Override:
  - Organization: cloud-team
    repo_name: repo1
    Repository:
      delete_branch_on_merge: false

```

### You want to check the rules for two organizations
```yaml
Organizations:
  - cloud-team
  - tech-ops

OrganizationStandards:
 - Organization: cloud-team
   Repository:
      delete_branch_on_merge: true

 - Organization: tech-ops
   Repository:
      delete_branch_on_merge: false
```

