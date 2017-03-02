package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo [name]",
	Short: "List repositories for user",
	RunE:  fetchRepos,
}

var contributorsCmd = &cobra.Command{
	Use:   "contrib [repo]",
	Short: "List the contributors for the repo",
	RunE:  fetchContributors,
}

func fetchRepos(cmd *cobra.Command, args []string) error {
	fmt.Println("Getting repositories data...")
	// var repos data.Repos
	// GHService.GetRepo(&repos)
	// sort.Sort(repos)
	// for i := 0; i < len(repos); i++ {
	// 	repos[i].Print()
	// }
	return nil
}

func fetchContributors(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		fmt.Println("Need to provide name of the repository")
		return nil
	}
	fmt.Println("Fetching list of contributors for repo: ", args[0])
	// repo := data.Repo{Name: args[0]}
	// var users data.Users
	// GHService.GetContributors(&repo, &users)
	// for i := 0; i < len(users); i++ {
	// 	users[i].Print()
	// }
	return nil
}

func init() {
	RootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(contributorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
