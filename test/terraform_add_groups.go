package test

import (
  "testing"
  "fmt"
  "encoding/json"

  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/stretchr/testify/assert"
)

func TestAddGroups(t *testing.T) {
  base := "../examples/add-groups"

  pathToGroupsData := fmt.Sprintf("%s/groups.auto.tfvars.json", base)

  expectedGitlabGroups := *GetGroupsData(pathToGroupsData)

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

  groupsOutAsObj := make(map[string]interface{})
  json.Unmarshal([]byte(groupsOut), &groupsOutAsObj)
  fmt.Println("Printing Groups")
  actualGitlabGroups := make(GitlabGroups)
  for key := range groupsOutAsObj {
    groupDetails := groupsOutAsObj[key].(map[string]interface{})
    group := GitlabGroup{
      GroupName(groupDetails["name"].(string)),
      groupDetails["description"].(string),
    }
    actualGitlabGroups[GroupName(key)] = group
  }

  assert.Equal(t, expectedGitlabGroups, actualGitlabGroups)
}
