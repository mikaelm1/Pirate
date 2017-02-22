package service

import "github.com/spf13/viper"

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
