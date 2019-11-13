output "gitlab_users" {
  value = var.users
}

output "gitlab_groups" {
  value = var.groups
}

output "memberships" {
  value = local.memberships
}
