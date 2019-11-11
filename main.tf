variable "base_url" {
  type = string
}

variable "users" {
  type = map
}

variable "groups" {
  type = map
}

variable "gitlab_token" {
  type = string
}

output "gitlab_users" {
  value = var.users
}

output "gitlab_token" {
  value = var.gitlab_token
}

output "gitlab_groups" {
  value = var.groups
}

locals {
  membershipAsArray = flatten([
    for email, user in var.users :
    [
      for group in split(":", user.groups) :
      {
        email = user.email
        group = group
      }
    ]
  ])
  memberships = { for membership in local.membershipAsArray : "${membership.email}_${membership.group}" => membership }
}

output "memberships" {
  value = local.memberships
}

provider "gitlab" {
  token    = var.gitlab_token
  base_url = var.base_url
  insecure = true
}

resource "gitlab_user" "user" {
  for_each          = var.users
  name              = each.value.username
  username          = each.value.username
  password          = each.value.password
  email             = each.value.email
  is_admin          = false
  projects_limit    = 100
  can_create_group  = true
  is_external       = false
  skip_confirmation = true
}

resource "gitlab_group" "group" {
  for_each    = var.groups
  name        = "${each.value.group_name}"
  path        = "${each.value.group_name}"
  description = "${each.value.group_description}"
}

resource "gitlab_group_membership" "group-membership" {
  for_each     = local.memberships
  group_id     = "${lookup(gitlab_group.group[format("%s", each.value.group)], "id", "0")}"
  user_id      = "${lookup(gitlab_user.user[format("%s", each.value.email)], "id", "0")}"
  access_level = "guest"
}
