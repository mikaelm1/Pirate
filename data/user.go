package data

import "fmt"

type User struct {
	DropletLimit    int    `json:"droplet_limit"`
	Email           string `json:"email"`
	UUID            string `json:"uuid"`
	FloatingIPLimit int    `json:"floating_ip_limit"`
	EmailVerified   bool   `json:"email_verified"`
	Status          string `json:"status"`
}

type Account struct {
	UserInfo User `json:"account"`
}

// PrintInfo displays an account's info
func (a *Account) PrintInfo() {
	fmt.Println("\n      **Your Account Info**")
	fmt.Println("Email:          ", a.UserInfo.Email)
	fmt.Println("Droplet Limit:  ", a.UserInfo.DropletLimit)
	fmt.Println("Email Verified: ", a.UserInfo.EmailVerified)
	fmt.Println()
}
