package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"encoding/json"

	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo [name]",
	Short: "List repositories for given user",
	RunE:  fetchRepos,
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"login"`
	CreatedAt string `json:"created_at"`
}

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

func fetchRepos(cmd *cobra.Command, args []string) error {
	fmt.Println("Getting repositories data for: " + args[0])
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v/repos", args[0]))
	if err != nil {
		fmt.Println("Error getting response: ", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	resp.Body.Read(body)
	var repos Repos
	json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	sort.Sort(repos)
	// output, err := json.MarshalIndent(&repo, "", " ")
	// os.Stdout.Write(output)
	for i := 0; i < len(repos); i++ {
		repos[i].Print()
	}
	return nil
}

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

func init() {
	RootCmd.AddCommand(repoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
