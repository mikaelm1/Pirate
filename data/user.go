package data

import "fmt"

// User model
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"login"`
	CreatedAt string `json:"created_at"`
}

// Users is an array of User structs
type Users []User

// Print displays info about user
func (u *User) Print() {
	fmt.Println("==================================================")
	fmt.Println("User: ", u.Username)
	fmt.Println("ID: ", u.ID)
}

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Username < u[j].Username
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
