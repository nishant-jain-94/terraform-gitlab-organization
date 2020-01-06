# Inputs

## Variables

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| base_url | Refers to the Url of the Gitlab Instance | string | N/A | Yes |
| gitlab_token | Refers to the personal token with Admin Privileges to create resources | string | N/A | Yes |
| users | List of users to be added onto Gitlab | map | {} | Yes |
| groups | List of groups to be created on Gitlab | map | {} | Yes |
| projects | List of projects to be created on Gitlab | map | {} | Yes |
| namespaces | List of the User and Group Namespaces used on Gitlab | map | {} | Yes |

## User

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| username | Refers to the username of the User | String | N/A | Yes |
| email | Refers to the email of the User | String | N/A | Yes |
| organization | Refers to the Organization of the User | String | N/A | Yes |
| password | Refers to the password of the User | String | N/A | Yes |
| groups_access | Refers to the Groups the User has access to. For more details checkout the [Group](#group) Section | Map | N/A | Yes |

## Project

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| name | Refers to the name of the Project | String | N/A | Yes |
| description | Refers to a succinct statement which describes the Project. | String | N/A | Yes |
| namespace_id | Refers to the name of the group under which the project needs to be created. | String | N/A | Yes |
| visibility_level | Refers to the visibility of the Project. It can have values `private`, `internal`, `public` | String | N/A | Yes
| only_allow_merge_if_pipeline_succeeds | Set to true if you want allow merges only if a pipeline succeeds. | Boolean String | N/A | Yes |
| shared_with_groups | Enable sharing the project with a set of groups. **Key**: Name of the Project. **Value**: Access level permission. Valid values are `guest`, `reporter`, `developer`, `master`  | Map | N/A | Yes |

## Group

A Key-Value pair. With key as the name of the group and the value containing the following properties

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| group_name | Refers to the name of the Group | String | N/A | Yes |
| group_description | Refers to a succinct statement used to describe the Group | String | N/A | Yes |
