package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mikaelm1/pirate/data"
)

// ReposWatched fetches the repos the current user watches
func (g *GitHubService) ReposWatched(repos *data.Repos) error {
	token := getUserToken()
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/user/subscriptions?access_token=%v", token))
	if err != nil {
		return err
	}
	// fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Read(body)
	err = json.Unmarshal(body, repos)
	if err != nil {
		return err
	}
	return nil
}

// IsWatchingRepo checks to see if user is currently watching the passed in repo
func (g *GitHubService) IsWatchingRepo(repo string, sub *data.Subscription) error {
	token := getUserToken()
	username := getUsername()
	res, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%v/%v/subscription?access_token=%v", username, repo, token))
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &sub)
	if err != nil {
		return err
	}
	return nil
}
