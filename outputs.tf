output "gitlab_users" {
  value = var.users
}

output "gitlab_groups" {
  value = var.groups
}

output "group_memberships" {
  value = local.group_memberships
}

output "project_memberships" {
  value = local.project_memberships
}
