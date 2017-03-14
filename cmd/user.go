package cmd

import "github.com/spf13/cobra"
import "fmt"

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Gets your account info",
	RunE:  fetchData,
}

func fetchData(cmd *cobra.Command, args []string) error {
	fmt.Println("Fetching your account info...")
	err := DOService.GetUserInfo()
	if err != nil {
		fmt.Println("There was an error fetching your info: ", err)
		return err
	}
	return nil
}

func init() {
	RootCmd.AddCommand(userCmd)
}
