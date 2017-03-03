// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/mikaelm1/pirate/data"
	"github.com/spf13/cobra"
)

var (
	listDropletsFlag bool
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
		fmt.Println(droplets)
		for i := 0; i < len(droplets.DropletsList); i++ {
			droplets.DropletsList[i].PrintInfo()
		}
	} else {
		fmt.Println("You must pass in one of the flags.")
	}
	return nil
}

func init() {
	RootCmd.AddCommand(dropletCmd)

	dropletCmd.Flags().BoolVarP(&listDropletsFlag, "list", "l", false, "list my droplets")

}
