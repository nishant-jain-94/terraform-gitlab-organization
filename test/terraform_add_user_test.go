package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type EmailId string
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

type GitlabUsers map[string]GitlabUser

type GitlabUsersFile struct {
	Users GitlabUsers `json:"users"`
}

// Gets all the user data from the tfvars json file
func GetUsersData() *GitlabUsers {
	jsonFile, err := os.Open("../examples/add-users/users.auto.tfvars.json")
	if err != nil {
		fmt.Println("Error in reading the file")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var gitlabUsersFile GitlabUsersFile
	json.Unmarshal([]byte(byteValue), &gitlabUsersFile)
	return &gitlabUsersFile.Users
}

// Tests if all Users in the auto varfile are added to the Gitlab instance under test
func TestAddGitlabUsers(t *testing.T) {
	t.Parallel()
	var expectedGitlabUsers GitlabUsers = *GetUsersData()
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/add-users",

		NoColor: true,
	}

	defer terraform.Destroy(t, terraformOptions)

	// Initialize and Apply Terraform
	terraform.InitAndApply(t, terraformOptions)
	out, err := terraform.RunTerraformCommandAndGetStdoutE(t, terraformOptions, "output", "-no-color", "-json", "gitlab_users")
	if err != nil {
		fmt.Println(err)
	}

	var actualGitlabUsers = make(GitlabUsers)
	json.Unmarshal([]byte(out), &actualGitlabUsers)

	// Replacing nil structs with empty objects to Groups and Projects Access
	for key := range actualGitlabUsers {
		var user = actualGitlabUsers[key]
		user.GroupsAccess = make(GroupsAccess)
		user.ProjectsAccess = make(ProjectsAccess)
		actualGitlabUsers[key] = user
	}

	// Test if the expected users are equal to the actual gitlab users
	assert.Equal(t, expectedGitlabUsers, actualGitlabUsers)
}
