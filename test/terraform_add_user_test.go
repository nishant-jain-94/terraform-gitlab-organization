package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Tests if all Users in the auto varfile are added to the Gitlab instance under test
func TestAddGitlabUsers(t *testing.T) {
  base := "../examples/add-users"
  pathToUsersData := fmt.Sprintf("%s/users.auto.tfvars.json", base)
	expectedGitlabUsers := *GetUsersData(pathToUsersData)
	terraformOptions := &terraform.Options{
		TerraformDir: base,

		NoColor: true,
	}

	defer terraform.Destroy(t, terraformOptions)

	// Initialize and Apply Terraform
	terraform.InitAndApply(t, terraformOptions)
	out, err := terraform.RunTerraformCommandAndGetStdoutE(t, terraformOptions, "output", "-no-color", "-json", "gitlab_users")
	if err != nil {
		fmt.Println(err)
	}

	actualGitlabUsers := make(GitlabUsers)
	json.Unmarshal([]byte(out), &actualGitlabUsers)

	// Replacing nil structs with empty objects to Groups and Projects Access
	for key := range actualGitlabUsers {
		user := actualGitlabUsers[key]
		user.GroupsAccess = make(GroupsAccess)
		user.ProjectsAccess = make(ProjectsAccess)
		actualGitlabUsers[key] = user
	}

	// Test if the expected users are equal to the actual gitlab users
	assert.Equal(t, expectedGitlabUsers, actualGitlabUsers)
}
