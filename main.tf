terraform {
  required_version = "~> 0.12"
  required_providers {
    gitlab = "~> 2.0"
  }
}

provider "gitlab" {
  token    = var.gitlab_token
  base_url = var.base_url
  insecure = true
}

# Creates a cohort of Gitlab Users onto this Gitlab Instance
resource "gitlab_user" "this" {
  for_each = var.users

  name              = each.value.username
  username          = each.value.username
  password          = each.value.password
  email             = each.value.email
  projects_limit    = 100
  is_external       = false
  skip_confirmation = true
}

# Creates a cohort of Gitlab Groups onto this Gitlab Instance
resource "gitlab_group" "this" {
  for_each = var.groups

  name        = each.value.group_name
  path        = each.value.group_name
  description = each.value.group_description
}

# Creates a new project under a given namespace
resource "gitlab_project" "this" {
  for_each = var.projects

  name                                  = each.value.name
  path                                  = each.value.name
  namespace_id                          = each.value.namespace_id
  description                           = each.value.description
  only_allow_merge_if_pipeline_succeeds = true
  visibility_level                      = each.value.visibility_level

  depends_on = [
    gitlab_group.this,
    gitlab_user.this
  ]
}

# Establishes membership of the Gitlab Users with Gitlab Groups
resource "gitlab_group_membership" "this" {
  for_each = local.group_memberships

  group_id     = lookup(gitlab_group.this[format("%s", each.value.group)], "id", "0")
  user_id      = lookup(gitlab_user.this[format("%s", each.value.email)], "id", "0")
  access_level = "guest"

  depends_on = [
    gitlab_user.this,
    gitlab_group.this
  ]
}

# resource "gitlab_project_share_group" "this" {
#   for_each = local.project_memberships

#   project_id =
#   group_id =
#   access_level =

#   depends_on = [
#     gitlab_group.this
#     gitlab_project.this
#   ]
# }
