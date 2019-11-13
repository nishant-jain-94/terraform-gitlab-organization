# tf-gitlab-provisioner-config

Configuration of Provisioning Gitlab Instance using Terraform

## Instructions

```
terraform init
terraform apply
```

## Inputs
| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| gitlab_url | Refers to the Url of the Gitlab Instance with the API Part | string | N/A | Yes |
| gitlab_token | Refers to the personal token with Admin Privileges | string | N/A | Yes |
| users | List of users to be added onto Gitlab | map | {} | No |
| groups | List of groups to be created on Gitlab | map | {} | No |

#### users

```terraform
{
  "john.doe@gmail.com" => {
    username = "john.doe"
    password = "password@123"
    email = "john.doe@gmail.com"
  }
}
```


#### groups

```terraform
{
  "reviewers" => {
    group_name = "reviewers"
    group_description = "List of Reviewers"
  }
}
```
