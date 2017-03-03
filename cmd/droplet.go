package cmd

import (
	"fmt"

	"errors"

	"github.com/mikaelm1/pirate/data"
	"github.com/spf13/cobra"
)

var (
	listDropletsFlag bool
	getSingleDroplet int
)

// dropletCmd represents the droplet command
var dropletCmd = &cobra.Command{
	Use:   "droplet",
	Short: "Run actions related to droplets.",
	RunE:  handleCommand,
}

func handleCommand(*cobra.Command, []string) error {
	if listDropletsFlag {
		fmt.Println("Fetching your droplets...")
		var droplets data.Droplets
		err := DOService.FetchDroplets(&droplets)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		for i := 0; i < len(droplets.DropletsList); i++ {
			droplets.DropletsList[i].PrintInfo()
		}
	} else if getSingleDroplet != -1 {
		return fetchDroplet()
	} else {
		return errors.New("flag missing")
	}
	return nil
}

func fetchDroplet() error {
	fmt.Println("Fetching droplet with id", getSingleDroplet)
	var droplet data.SingleDroplet
	err := DOService.FetchDroplet(&droplet, getSingleDroplet)
	if err != nil {
		return err
	}
	droplet.SDroplet.PrintInfo()
	return nil
}

func init() {
	RootCmd.AddCommand(dropletCmd)

	dropletCmd.Flags().BoolVarP(&listDropletsFlag, "list", "l", false, "list my droplets")
	dropletCmd.Flags().IntVarP(&getSingleDroplet, "single", "s", -1, "get droplet by id")

}
