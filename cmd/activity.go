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
	"sort"

	"github.com/mikaelm1/pirate/data"
	"github.com/spf13/cobra"
)

var (
	reposWatched string
)

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Commands related to user's activites on GitHub",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: performActivity,
}

func performActivity(cmd *cobra.Command, args []string) {
	if reposWatched != "" {
		fetchWatched()
	} else {
		fmt.Println("Choose a flag to fetch info about one of your activites")
	}
}

func fetchWatched() {
	fmt.Println("Fetching watched")
	var repos data.Repos
	err := GHService.ReposWatched(&repos)
	if err != nil {
		fmt.Println("There was an error getting your watched repos:")
		fmt.Println(err)
		return
	}
	sort.Sort(repos)
	for i := 0; i < len(repos); i++ {
		repos[i].Print()
	}
}

func init() {
	RootCmd.AddCommand(activityCmd)

	activityCmd.Flags().StringVarP(&reposWatched, "watched", "w", "", "Fetches the repos you are currently watching")
	activityCmd.Flags().Lookup("watched").NoOptDefVal = "true"

}
