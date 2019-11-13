# Transforms all the users and establishes membership with the Gitlab Group.
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
