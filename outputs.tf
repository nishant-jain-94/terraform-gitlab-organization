output "gitlab_users" {
  value = gitlab_user.this
}

output "gitlab_groups" {
  value = gitlab_group.this
}

output "group_memberships" {
  value = local.group_memberships
}

output "projects" {
  value = gitlab_project.this
}
