output "gitlab_users" {
  value = gitlab_user.this
}

output "gitlab_groups" {
  value = gitlab_group.this
}

output "group_memberships" {
  value = local.group_memberships
}

output "user" {
  value = lookup(gitlab_user.this["aditya.singh@northwind.in"], "id")
}

output "projects" {
  value = gitlab_project.this
}
