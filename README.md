# WIP: tf-gitlab-configuration

Configuration of Provisioning Gitlab Instance using Terraform

## Instructions

```
terraform init
terraform apply
```

## Inputs
| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| gitlab_url | Refers to the Url of the Gitlab Instance with the API Part | string | N/A | Yes |
| gitlab_token | Refers to the personal token with Admin Privileges | string | N/A | Yes |
| users | List of users to be added onto Gitlab | map | {} | Yes |
| groups | List of groups to be created on Gitlab | map | {} | Yes |
| projects | List pf projects to be created on Gitlab | map | {} | Yes |

#### users

```terraform
{
  users => {
    "john.doe@northwind.in" => {
      username = "john.doe"
      email = "john.doe@northwind.in"
      organization = "northwind"
      password = "password@123"
      groups_access = {
        northwind-wave-1 = "guest"
      },
      projects_access = {
        users-project = "guest
      }
    }
  }
}
```


#### groups

```terraform
{
  "groups" => {
    "northwind-wave-1" => {
      group_name = "reviewers"
      group_description = "List of Reviewers"
    }
  }
}
```

#### projects

```terraform
{
  "projects" => {
    mentors-projects = {
      name = "mentors-project"
      description = "This the Test Project"
      namespace_id = "northwind-mentors"
      visibility_level = "private"
      shared_with_groups = {
        northwind-auditors = "guest"
      }
    }
  }
}
```
