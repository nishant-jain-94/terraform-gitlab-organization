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
  default = {}
}

variable "group_namespaces" {
  default = {}
}
