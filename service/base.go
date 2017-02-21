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
	json.Unmarshal(body, repos)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
