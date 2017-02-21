package data

import "fmt"

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"login"`
	CreatedAt string `json:"created_at"`
}

type Users []User

type Repo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Owner       User   `json:"owner"`
	Private     bool   `json:"private"`
	Description string `json:"description"`
	CrearedAt   string `json:"created_at"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	BranchesURL string `json:"branches_url"`
	IssuesURL   string `json:"issues_url"`
}

type Repos []Repo

// Print prints the repo data to stdout
func (repo *Repo) Print() {
	fmt.Println("==================================================")
	fmt.Printf("Repo Name: %v\n", repo.Name)
	fmt.Printf("Private: %v\n", repo.Private)
	fmt.Printf("Created At: %v\n", repo.CrearedAt)
	fmt.Printf("Description: %v\n", repo.Description)
	fmt.Printf("Owner: %v\n", repo.Owner.Username)
}

func (repos Repos) Len() int {
	return len(repos)
}

func (repos Repos) Less(i, j int) bool {
	return repos[i].CrearedAt < repos[j].CrearedAt
}

func (repos Repos) Swap(i, j int) {
	repos[i], repos[j] = repos[j], repos[i]
}

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
