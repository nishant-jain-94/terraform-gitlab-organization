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

output "gitlab_users" {
  value = module.add_users_projects.gitlab_users
}

output "gitlab_projects" {
  value = module.add_users_projects.gitlab_projects
}

output "gitlab_groups" {
  value = module.add_users_projects.gitlab_groups
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
