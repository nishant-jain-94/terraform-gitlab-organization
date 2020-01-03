package test

import (
  "testing"
  "fmt"
  "encoding/json"

  "github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAddGroupUsers(t *testing.T) {
  base := "../examples/add-group-users"

  pathToGroupsData := fmt.Sprintf("%s/groups.auto.tfvars.json", base)
  pathToUsersData := fmt.Sprintf("%s/users.auto.tfvars.json", base)

  expectedGitlabGroups := *GetGroupsData(pathToGroupsData)
  expectedGitlabUsers := *GetUsersData(pathToUsersData)

  terraformOptions := &terraform.Options{
    TerraformDir: base,

    NoColor: true,
  }

  defer terraform.Destroy(t, terraformOptions)

  terraform.InitAndApply(t, terraformOptions)

  groupsOut, err := terraform.RunTerraformCommandAndGetStdoutE(t, terraformOptions, "output", "-no-color", "-json", "gitlab_groups")
  if err != nil {
    fmt.Println(err)
  }

  usersOut, err := terraform.RunTerraformCommandAndGetStdoutE(t, terraformOptions, "output", "-no-color", "-json", "gitlab_users")
  if err != nil {
    fmt.Println(err)
  }

  groupsOutAsObj := make(map[string]interface{})
  json.Unmarshal([]byte(groupsOut), &groupsOutAsObj)
  actualGitlabGroups := make(GitlabGroups)
  for key := range groupsOutAsObj {
    groupDetails := groupsOutAsObj[key].(map[string]interface{})
    group := GitlabGroup{
      GroupName(groupDetails["name"].(string)),
      groupDetails["description"].(string),
    }
    actualGitlabGroups[GroupName(key)] = group
  }

  actualGitlabUsers := make(GitlabUsers)
	json.Unmarshal([]byte(usersOut), &actualGitlabUsers)

	// Replacing nil structs with empty objects to Groups and Projects Access
	for key := range actualGitlabUsers {
		user := actualGitlabUsers[key]
		user.GroupsAccess = make(GroupsAccess)
		user.ProjectsAccess = make(ProjectsAccess)
		actualGitlabUsers[key] = user
	}

  assert.Equal(t, expectedGitlabUsers, actualGitlabUsers)
  assert.Equal(t, expectedGitlabGroups, actualGitlabGroups)
}
