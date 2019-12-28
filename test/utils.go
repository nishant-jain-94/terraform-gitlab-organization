package test

import (
  "fmt"
  "os"
  "io/ioutil"
  "encoding/json"
)

func readJsonFile(path string) []byte {
  jsonFile, err := os.Open(path)
  if err != nil {
    fmt.Println("Error in reading the file")
    fmt.Println(err)
  }
  defer jsonFile.Close()
  byteValue, _ := ioutil.ReadAll(jsonFile)
  return []byte(byteValue)
}

// Gets all the user data from the tfvars json file
func GetUsersData(path string) *GitlabUsers {
  usersData := readJsonFile(path)
  var gitlabUsersFile GitlabUsersFile
  json.Unmarshal(usersData, &gitlabUsersFile)
	return &gitlabUsersFile.Users
}

func GetProjectsData(path string) *GitlabProjects {
  projectsData := readJsonFile(path)
  var gitlabProjectsFile GitlabProjectsFile
  json.Unmarshal(projectsData, &gitlabProjectsFile)
  return &gitlabProjectsFile.Projects
}

func GetGroupsData(path string) *GitlabGroups {
  groupsData := readJsonFile(path)
  var gitlabGroupsFile GitlabGroupsFile
  json.Unmarshal(groupsData, &gitlabGroupsFile)
  return &gitlabGroupsFile.Groups
}
