variable "base_url" {
  type = "string"
}

variable "gitlab_token" {
  type = "string"
}

variable "users" {
  description = "Refers to all the users to be provisioned onto Gitlab"
  default     = {}
}

variable "groups" {
  description = "Refers to all the groups to be provisioned onto Gitlab"
  default     = {}
}

variable "projects" {
  description = "Refers to all the projects to be provisioned onto Gitlab"
  default     = {}
}

variable "user_namespaces" {
  description = "Refers to all the users which will be referenced in the groups and projects vars files. This helps populating data upfront for all the user resources which have been created in and out of the scope of this Terraform Script but needs to be referenced in Terraform Script."
  default     = {}
}

variable "group_namespaces" {
  description = "Refers to all the groups which will be referenced in the groups, users and project vars files. This helps populating data upfront for all the group resouces which have been created in and out of the scope of this Terraform Script but needs to referenced in Terraform Script."
  default     = {}
}
