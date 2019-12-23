module "add-users" {
  source = "../../"
  users = {
    "john.doe@northwind.com" = {
      username        = "john.doe"
      email           = "john.doe@gmail.com"
      password        = "password@123"
      organization    = "northwind"
      groups_access   = {}
      projects_access = {}
    }
  }
}
