package test

type EmailId string
type UserName string
type GroupName string
type GroupAccessLevel string
type ProjectName string
type ProjectAccessLevel string
type GroupsAccess map[GroupName]GroupAccessLevel
type ProjectsAccess map[ProjectName]ProjectAccessLevel

type GitlabUser struct {
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	GroupsAccess   GroupsAccess   `json:"groups_access"`
	ProjectsAccess ProjectsAccess `json:"projects_access"`
}

type GitlabProject struct {
  Name              ProjectName   `json:"name"`
  Description       string        `json:"description"`
  VisibilityLevel   string        `json:"visibility_level"`
}

type GitlabGroup struct {
  GroupName         GroupName `json:"group_name"`
  GroupDescription  string    `json:"group_description"`
}

type GitlabUsers map[UserName]GitlabUser
type GitlabGroups map[GroupName]GitlabGroup
type GitlabProjects map[ProjectName]GitlabProject

type GitlabUsersFile struct {
	Users GitlabUsers `json:"users"`
}

type GitlabGroupsFile struct {
  Groups GitlabGroups `json:"groups"`
}

type GitlabProjectsFile struct {
  Projects GitlabProjects `json:"projects"`
}
