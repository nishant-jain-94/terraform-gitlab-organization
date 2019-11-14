# Transforms all the users and establishes membership with the Gitlab Group and Projects.
locals {
  group_memberships_as_array = flatten([
    for email, user in var.users :
    [
      for group in split(":", user.groups) :
      {
        email = user.email
        group = group
      }
    ]
  ])

  group_memberships = {
    for group_membership in local.group_memberships_as_array :
    "${group_membership.email}_${group_membership.group}" => group_membership
  }

  project_memberships_as_array = flatten([
    for project_name, project in var.projects :
    [
      for group in split(":", project.groups) :
      {
        project_name = project.project_name
        group_name   = project.group_name
      }
    ]
  ])

  project_memberships = {
    for project_membership in local.project_memberships_as_array :
    "${project_membership.project_name}_${project_membership.group}" => project_membership
  }
}
