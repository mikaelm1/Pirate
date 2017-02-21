package cmd

import (
	"fmt"
	"sort"

	"github.com/mikaelm1/pirate/data"
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo [name]",
	Short: "List repositories for given user",
	RunE:  fetchRepos,
}

func fetchRepos(cmd *cobra.Command, args []string) error {
	fmt.Println("Getting repositories data for: " + args[0])
	var repos data.Repos
	GHService.GetRepo(&repos)
	sort.Sort(repos)
	for i := 0; i < len(repos); i++ {
		repos[i].Print()
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
