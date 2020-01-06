# terraform-gitlab-organization

Configuration of Provisioning Gitlab Instance using Terraform


### <details><summary>Add Users</summary>
<p>

####  main.tf
```hcl
variable "gitlab_token" {
  type = "string"
}

variable "base_url" {
  type = "string"
}

variable "users" {
  default = {}
}

module "add_users" {
  source       = "../../"
  gitlab_token = var.gitlab_token
  base_url     = var.base_url
  users        = var.users
}
```

#### users.auto.tfvars.json

```json
{
  "users": {
    "aditya.singh@northwind.in": {
      "username": "aditya.singh",
      "email": "aditya.singh@northwind.in",
      "password": "password@123",
      "groups_access": {},
      "projects_access": {}
    },
    "raj.singh@northwind.in": {
      "username": "raj.singh",
      "email": "raj.singh@northwind.in",
      "password": "password@123",
      "groups_access": {},
      "projects_access": {}
    }
  }
}
```
</p>
</details>

### Add Groups

#### main.tf

```hcl
variable "gitlab_token" {
  type = "string"
}

variable "base_url" {
  type = "string"
}

variable groups {
  default = {}
}

module "add_groups" {
  source       = "../../"
  gitlab_token = var.gitlab_token
  base_url     = var.base_url
  groups       = var.groups
}
```

#### groups.auto.tfvars.json
```json
{
  "groups": {
    "northwind-wave-2": {
      "group_name": "northwind-wave-2",
      "group_description": "A Group for entire northwind Wave 1"
    },
    "northwind-mentors-1": {
      "group_name": "northwind-mentors-1",
      "group_description": "Group Created for northwind mentors"
    },
    "northwind-auditors-1": {
      "group_name": "northwind-auditors-1",
      "group_description": "Group Created for auditors"
    }
  }
}
```

### Add Users to Groups
#### main.tf
```hcl
variable "gitlab_token" {
  type = "string"
}

variable "base_url" {
  type = "string"
}

variable "users" {
  default = {}
}

variable "groups" {
  default = {}
}

variable "user_namespaces" {
  default = {}
}

variable "group_namespaces" {
  default = {}
}

module "add_groups_users" {
  source           = "../../"
  gitlab_token     = var.gitlab_token
  base_url         = var.base_url
  groups           = var.groups
  users            = var.users
  user_namespaces  = var.user_namespaces
  group_namespaces = var.group_namespaces
}
```

#### groups.auto.tfvars.json
```json
{
  "groups": {
    "northwind-wave-2": {
      "group_name": "northwind-wave-2",
      "group_description": "A Group for entire northwind Wave 1"
    },
    "northwind-mentors-1": {
      "group_name": "northwind-mentors-1",
      "group_description": "Group Created for northwind mentors"
    },
    "northwind-auditors-1": {
      "group_name": "northwind-auditors-1",
      "group_description": "Group Created for auditors"
    }
  }
}
```

#### users.auto.tfvars.json
```json
{
  "users": {
    "aditya.singh@northwind.in": {
      "username": "aditya.singh",
      "email": "aditya.singh@northwind.in",
      "password": "password@123",
      "groups_access": {},
      "projects_access": {}
    },
    "raj.singh@northwind.in": {
      "username": "raj.singh",
      "email": "raj.singh@northwind.in",
      "password": "password@123",
      "groups_access": {},
      "projects_access": {}
    }
  }
}
```

#### namespaces.auto.tfvars.json

```json
{
  "group_namespaces": {
    "northwind-wave-2": "northwind-wave-2",
    "northwind-mentors-1": "northwind-mentors-1",
    "northwind-auditors-1": "northwind-auditors-1"
  },
  "user_namespaces": {
    "aditya.singh": "aditya.singh@northwind.in",
    "raj.singh": "raj.singh@northwind.in",
    "sachin.grover": "sachin@northwind.in",
    "sagar.patke": "sagar.patke@northwind.in"
  }
}
```

### Add Projects to Groups

#### main.tf

```hcl
variable "gitlab_token" {
  type = "string"
}

variable "base_url" {
  type = "string"
}

variable "users" {
  default = {}
}

variable "projects" {
  default = {}
}

variable "groups" {
  default = {}
}

variable "user_namespaces" {
  default = {}
}

variable "group_namespaces" {
  default = {}
}

module "add_users_projects" {
  source           = "../../"
  gitlab_token     = var.gitlab_token
  base_url         = var.base_url
  projects         = var.projects
  groups           = var.groups
  user_namespaces  = var.user_namespaces
  group_namespaces = var.group_namespaces
}
```

#### groups.auto.tfvars.json

```json
{
  "groups": {
    "northwind-wave-2": {
      "group_name": "northwind-wave-2",
      "group_description": "A Group for entire northwind Wave 1"
    },
    "northwind-mentors-1": {
      "group_name": "northwind-mentors-1",
      "group_description": "Group Created for northwind mentors"
    },
    "northwind-auditors-1": {
      "group_name": "northwind-auditors-1",
      "group_description": "Group Created for auditors"
    }
  }
}
```

#### namespaces.auto.tfvars.json
```json
{
  "group_namespaces": {
    "northwind-wave-2": "northwind-wave-2",
    "northwind-mentors-1": "northwind-mentors-1",
    "northwind-auditors-1": "northwind-auditors-1"
  },
  "user_namespaces": {}
}
```

#### projects.auto.tfvars.json
```json
{
  "projects": {
    "mentors-project-1": {
      "name": "mentors-project-1",
      "description": "This is only the Test Project",
      "visibility_level": "private",
      "namespace_id": "northwind-mentors-1",
      "only_allow_merge_if_pipeline_succeeds": false,
      "shared_with_groups": {}
    },
    "users-project-1": {
      "name": "users-project-1",
      "description": "This is only the Test Project",
      "visibility_level": "private",
      "namespace_id": "northwind-auditors-1",
      "only_allow_merge_if_pipeline_succeeds": true,
      "shared_with_groups": {
        "northwind-wave-2": "guest"
      }
    }
  }
}
```

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
