package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mikaelm1/pirate/data"

	"net/http"

	"github.com/spf13/viper"
)

type GitHubService struct {
	// Client *http.Client
}

type accessToken struct {
	token string `yaml:"token"`
}

func getUserToken() string {
	tkn := viper.GetString("token")
	return tkn
}

func getUsername() string {
	username := viper.GetString("username")
	return username
}

//GetRepo fetches all the repos for the authenticated user
func (c *GitHubService) GetRepo(repos *data.Repos) {
	token := getUserToken()
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/user/repos?access_token=%v", token))
	if err != nil {
		fmt.Println("Error getting response: ", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	resp.Body.Read(body)
	err = json.Unmarshal(body, repos)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// GetContributors fetches all the contributos for the given repo
func (c *GitHubService) GetContributors(repo *data.Repo, users *data.Users) error {
	token := getUserToken()
	user := getUsername()
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%v/%v/contributors?access_token=%v", user, repo.Name, token))
	if err != nil {
		fmt.Println("Error getting response: ", err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	resp.Body.Read(body)
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Error unmarshaling users")
	}
	return nil
}
