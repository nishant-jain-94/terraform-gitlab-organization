package test

import (
  "testing"
  "fmt"
  "encoding/json"

  "github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAddUsersProjects(t *testing.T) {
  base := "../examples/add-group-projects"

  pathToProjectsData := fmt.Sprintf("%s/projects.auto.tfvars.json", base)
  pathToGroupsData := fmt.Sprintf("%s/groups.auto.tfvars.json", base)

  expectedGitlabProjects := *GetProjectsData(pathToProjectsData)
  expectedGitlabGroups := *GetGroupsData(pathToGroupsData)

  terraformOptions := &terraform.Options{
    TerraformDir: base,

    NoColor: true,
  }

  defer terraform.Destroy(t, terraformOptions)

  terraform.InitAndApply(t, terraformOptions)
  projectsOut, err := terraform.RunTerraformCommandAndGetStdoutE(t, terraformOptions, "output", "-no-color", "-json", "gitlab_projects")
  if err != nil {
    fmt.Println(err)
  }

  groupsOut, err := terraform.RunTerraformCommandAndGetStdoutE(t, terraformOptions, "output", "-no-color", "-json", "gitlab_groups")
  if err != nil {
    fmt.Println(err)
  }

  groupsOutAsObj := make(map[string]interface{})
  json.Unmarshal([]byte(groupsOut), &groupsOutAsObj)
  fmt.Println("Printing Groups")
  fmt.Println(groupsOutAsObj)
  actualGitlabGroups := make(GitlabGroups)
  for key := range groupsOutAsObj {
    groupDetails := groupsOutAsObj[key].(map[string]interface{})
    group := GitlabGroup{
      GroupName(groupDetails["name"].(string)),
      groupDetails["description"].(string),
    }
    actualGitlabGroups[GroupName(key)] = group
  }

  fmt.Println(actualGitlabGroups)
  actualGitlabProjects := make(GitlabProjects)
  json.Unmarshal([]byte(projectsOut), &actualGitlabProjects)

  assert.Equal(t, expectedGitlabGroups, actualGitlabGroups)
  assert.Equal(t, expectedGitlabProjects, actualGitlabProjects)
}
