variable "gitlab_token" {
  type = "string"
}

variable "base_url" {
  type = "string"
}

variable "users" {
  default = {}
}

output "gitlab_users" {
  value = module.add_users.gitlab_users
}

output "gitlab_groups" {
  value = module.add_users.gitlab_groups
}

output "group_memberships" {
  value = module.add_users.group_memberships
}

output "projects" {
  value = module.add_users.projects
}

module "add_users" {
  source       = "../../"
  gitlab_token = var.gitlab_token
  base_url     = var.base_url
  users        = var.users
}
