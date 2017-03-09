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

	"errors"

	"github.com/mikaelm1/pirate/data"
	"github.com/spf13/cobra"
)

var (
	listSSHKeys bool
)

// sshkeyCmd represents the sshkey command
var sshkeyCmd = &cobra.Command{
	Use:   "ssh_key",
	Short: "Commands related to your ssh keys",
	RunE:  handleSSHCommand,
}

func handleSSHCommand(*cobra.Command, []string) error {
	if listSSHKeys {
		return getAllSSHKeys()
	}
	return errors.New("Need to provide a flag")
}

func getAllSSHKeys() error {
	fmt.Println("Fetching your ssh keys...")
	var keys data.ArraySSHKey
	_, err := DOService.FetchAllSSHKeys(&keys)
	if err != nil {
		return err
	}
	if len(keys.SSHKey) > 0 {
		keys.PrintInfo()
	} else {
		fmt.Println("You do not have any keys")
	}
	return nil
}

func init() {
	RootCmd.AddCommand(sshkeyCmd)

	sshkeyCmd.Flags().BoolVarP(&listSSHKeys, "list", "l", false, "get all of your keys")

}
