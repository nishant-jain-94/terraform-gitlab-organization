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

output "gitlab_groups" {
  value = module.add_groups_users.gitlab_groups
}

output "gitlab_users" {
  value = module.add_groups_users.gitlab_users
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
