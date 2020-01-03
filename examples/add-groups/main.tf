variable "gitlab_token" {
  type = "string"
}

variable "base_url" {
  type = "string"
}

variable groups {
  default = {}
}

output "gitlab_groups" {
  value = module.add_groups.gitlab_groups
}

module "add_groups" {
  source       = "../../"
  gitlab_token = var.gitlab_token
  base_url     = var.base_url
  groups       = var.groups
}
