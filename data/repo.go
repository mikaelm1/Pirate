package data

import "fmt"

// Repo model
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

// Repos is an array of Repo objects
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
