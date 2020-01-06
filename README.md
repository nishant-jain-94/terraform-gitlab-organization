# terraform-gitlab-organization ![](https://github.com/nishant-jain-94/terraform-gitlab-organization/workflows/terraform_testing/badge.svg)

A Terraform Module to Provision Gitlab Resources on Gitlab Instance

## :rocket: Features

- [Add Users](#add-users)
- [Add Groups](#add-groups)
- [Add Users to Groups](#add-users-to-groups)
- [Add Projects to Groups](#add-projects-to-groups)

## :information_desk_person: Usage

### Add Users

<details>
<summary>Add Users Details</summary>

####  main.tf
```terraform
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
</details>

### Add Groups

<details>
<summary>Add Groups</summary>

#### main.tf

```terraform
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
</details>

### Add Users to Groups

<details>
<summary>Add Users to Groups</summary>

#### main.tf
```terraform
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
</details>

### Add Projects to Groups

<details>
<summary>Add Projects to Groups</summary>
#### main.tf

```terraform
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

</details>


## Contribution Guidelines

Contributions to this Module are very much welcome. Checkout the [Contribution Guidelines](./CONTRIBUTING.md) for instructions.

## Changelog

Changelog to this terraform module are logged here on [Changelog](./CHANGELOG.md).

## Issues



## License
This code is release under MIT License. Please see [LICENSE](./LICENSE) for more details.
