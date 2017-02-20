package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo [name]",
	Short: "List repositories for given user",
	RunE:  fetchRepos,
}

type Repo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

func fetchRepos(cmd *cobra.Command, args []string) error {
	fmt.Println("Called repo for user: " + args[0])
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v/repos", args[0]))
	if err != nil {
		fmt.Println("Error getting response: ", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	resp.Body.Read(body)
	var repo []Repo
	json.Unmarshal(body, &repo)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	// output, err := json.MarshalIndent(&repo, "", " ")
	// os.Stdout.Write(output)
	for i := 0; i < len(repo); i++ {
		fmt.Printf("Repo Name: %v\nPrivate: %v\n", repo[i].Name, repo[i].Private)
		fmt.Println("===========================")
	}
	return nil
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
