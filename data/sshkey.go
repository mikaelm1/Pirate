package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// SSHKey model
type SSHKey struct {
	ID          int    `json:"id"`
	Fingerprint string `json:"fingerprint"`
	PublicKey   string `json:"public_key"`
	Name        string `json:"name"`
}

// SingleSSHKey is model for unmarhsalling a single ssh key
type SingleSSHKey struct {
	SSHKey *SSHKey `json:"ssh_key"`
}

// ArraySSHKey is model for unmarhsalling an array of ssh keys
type ArraySSHKey struct {
	SSHKey []SSHKey `json:"ssh_keys"`
}

// PrintInfo displays info about the array of keys
func (arr *ArraySSHKey) PrintInfo() {
	for _, v := range arr.SSHKey {
		v.PrintInfo()
	}
}

// PrintInfo displays info about a single key
func (k *SSHKey) PrintInfo() {
	if viper.GetString("output") == "json" {
		k.JSONPrint()
	} else {
		k.TextPrint()
	}
}

// TextPrint displays info in text format
func (k *SSHKey) TextPrint() {
	fmt.Println("==============================")
	fmt.Println("ID:           ", k.ID)
	fmt.Println("Fingerprint:  ", k.Fingerprint)
	fmt.Println("Name:         ", k.Name)
}

// JSONPrint displays info in JSON format
func (k *SSHKey) JSONPrint() {
	output, err := json.MarshalIndent(k, "", "    ")
	if err != nil {
		fmt.Println("Error parsing to JSON")
	}
	os.Stdout.Write(output)
}
