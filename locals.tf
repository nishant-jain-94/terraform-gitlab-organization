# Transforms all the users and establishes membership with the Gitlab Group and Projects.
locals {
  group_memberships_as_array = flatten([
    for email, user in var.users :
    [
      for group_name, access_level in user.groups_access :
      {
        user_email   = user.email
        group_name   = group_name
        access_level = access_level
      }
    ]
  ])

  group_memberships = {
    for group_membership in local.group_memberships_as_array :
    "${group_membership.user_email}_${group_membership.group_name}" => group_membership
  }

  project_user_membership_as_array = flatten([
    for user_email, user in var.users :
    [
      for project_name, access_level in user.projects_access :
      {
        project_name = project_name
        user_email   = user_email
        access_level = access_level
      }
    ]
  ])

  project_user_memberships = {
    for project_membership in local.project_user_membership_as_array :
    "${project_membership.project_name}_${project_membership.user_email}" => project_membership
  }
}
