package service

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type GitHubService struct {
	// Client *http.Client
}

type DOService struct {
	client *http.Client
}

func (d *DOService) Client() {
	d.client = &http.Client{
		Timeout: time.Second * 30,
	}
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
