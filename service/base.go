package service

import (
	"fmt"
	"net/http"
	"time"

	"bytes"

	"github.com/spf13/viper"
)

type GitHubService struct {
	// Client *http.Client
}

type DOService struct {
	client *http.Client
}

// Client initializes the http client
func (c *DOService) Client() {
	c.client = &http.Client{
		Timeout: time.Second * 30,
	}
}

// MakeGETRequest is a generic helper function for making GET requests
func (c *DOService) MakeGETRequest(url string) (*http.Response, error) {
	token := getUserToken()
	bearer := fmt.Sprintf("Bearer %v", token)
	c.Client()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", bearer)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MakePostRequest is a generic helper function for making POSST requests
func (c *DOService) MakePostRequest(url string, body []byte) (*http.Response, error) {
	token := getUserToken()
	bearer := fmt.Sprintf("Bearer %v", token)
	c.Client()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
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
