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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}