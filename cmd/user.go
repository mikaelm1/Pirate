package cmd

import (
	"fmt"

	"github.com/mikaelm1/pirate/data"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Gets your account info",
	RunE:  fetchData,
}

func fetchData(cmd *cobra.Command, args []string) error {
	fmt.Println("Fetching your account info...")
	var account data.Account
	err := DOService.GetUserInfo(&account)
	if err != nil {
		fmt.Println("There was an error fetching your info: ", err)
		return err
	}
	account.PrintInfo()
	return nil
}

func init() {
	RootCmd.AddCommand(userCmd)
}
