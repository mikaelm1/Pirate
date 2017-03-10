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

	"io/ioutil"

	"github.com/mikaelm1/pirate/data"
	"github.com/spf13/cobra"
)

var (
	listSSHKeys      bool
	showSingleKey    bool
	sshKeyID         int
	sshKeyFingerpint string
	publicKeyPath    string
	publicKeyString  string
	keyName          string
	deleteKeyFinger  string
	deleteKeyID      string
)

// sshkeyCmd represents the sshkey command
var sshkeyCmd = &cobra.Command{
	Use:   "ssh_key",
	Short: "Commands related to your ssh keys",
	RunE:  handleSSHCommand,
}

var sshKeyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new ssh key",
	RunE:  createNewSSHKey,
}

var sshKeyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an ssh key",
	RunE:  deleteSSHKey,
}

func deleteSSHKey(*cobra.Command, []string) error {
	if deleteKeyFinger == "" && deleteKeyID == "" {
		return errors.New("Must provide either the ID or Fingerpint of ssh key to delete")
	}
	var _id string
	if deleteKeyFinger != "" {
		_id = deleteKeyFinger
	} else {
		_id = deleteKeyID
	}
	fmt.Println("Deleting your ssh key...")
	_, err := DOService.DeleteSSHKey(_id)
	if err != nil {
		return err
	}
	fmt.Println("SSH key successfully deleted")
	return nil
}

func createNewSSHKey(*cobra.Command, []string) error {
	fmt.Println("Creating new SSH key...")
	if keyName == "" {
		return errors.New("New key must have a name")
	}
	if publicKeyPath == "" && publicKeyString == "" {
		return errors.New("Must provide either a path to your public key or the key itself")
	}
	var _key string
	if publicKeyPath != "" {
		dat, err := ioutil.ReadFile(publicKeyPath)
		if err != nil {
			return err
		}
		// fmt.Println(string(dat))
		_key = string(dat)
		// fmt.Println("Creating new key using public key:", string(dat))
	} else if publicKeyString != "" {
		_key = publicKeyString
		// fmt.Println("Creating new key using public key: ", publicKeyString)
	}
	key := data.SSHKey{
		Name:      keyName,
		PublicKey: _key,
	}
	var singleKey data.SingleSSHKey
	_, err := DOService.CreateSSHKey(&key, &singleKey)
	if err != nil {
		return err
	}
	fmt.Println("New SSH key created")
	if outputType == "json" {
		singleKey.SSHKey.JSONPrint()
	} else {
		singleKey.SSHKey.PrintInfo()
	}
	return nil
}

func handleSSHCommand(*cobra.Command, []string) error {
	if listSSHKeys {
		return getAllSSHKeys()
	} else if showSingleKey {
		return getSingleKey()
	}
	return errors.New("Must to provide a flag")
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

func getSingleKey() error {
	_id := ""
	if sshKeyID != 0 {
		fmt.Println("Fetching key with ID:", sshKeyID)
		_id = fmt.Sprintf("%d", sshKeyID)
	} else if sshKeyFingerpint != "" {
		_id = sshKeyFingerpint
		fmt.Println("Fetching key with Fingerpint:", sshKeyFingerpint)
	} else {
		return errors.New("Need to provide either an ID or Fingerpint")
	}
	var key data.SingleSSHKey
	_, err := DOService.FetchSingleKey(_id, &key)
	if err != nil {
		return err
	}
	if outputType == "json" {
		key.SSHKey.JSONPrint()
	} else {
		key.SSHKey.PrintInfo()
	}
	return nil
}

func init() {
	RootCmd.AddCommand(sshkeyCmd)
	sshkeyCmd.AddCommand(sshKeyCreateCmd)
	sshkeyCmd.AddCommand(sshKeyDeleteCmd)

	sshkeyCmd.Flags().BoolVarP(&listSSHKeys, "list", "l", false, "get all of your keys")
	sshkeyCmd.Flags().BoolVarP(&showSingleKey, "single", "s", false, "Get key with either Id or Fingerpint")
	sshkeyCmd.Flags().IntVar(&sshKeyID, "id", 0, "The ID of your ssh key")
	sshkeyCmd.Flags().StringVarP(&sshKeyFingerpint, "fingerprint", "f", "", "The Fingerprint of your ssh key")
	// create key Flags
	sshKeyCreateCmd.Flags().StringVarP(&publicKeyPath, "keypath", "k", "", "The absolute path to your public ssh key")
	sshKeyCreateCmd.Flags().StringVarP(&publicKeyString, "public-key", "p", "", "Your public ssh key")
	sshKeyCreateCmd.Flags().StringVarP(&keyName, "key-name", "n", "", "The name of the new ssh key")
	// delete key Flags
	sshKeyDeleteCmd.Flags().StringVarP(&deleteKeyID, "key-id", "i", "", "The ID of your ssh key")
	sshKeyDeleteCmd.Flags().StringVarP(&deleteKeyFinger, "fingerprint", "f", "", "The fingerprint of your ssh key")
}
