variable "base_url" {
  default = "https://gitlab.com"
}

variable "gitlab_token" {
  type = string
}

variable "users" {
  description = "Refers to all the users to be provisioned onto Gitlab"
  default     = {}
}

variable "groups" {
  description = "Refers to all the groups to be provisioned onto Gitlab"
  default     = {}
}
