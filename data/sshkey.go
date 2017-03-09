package data

import (
	"fmt"
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
	fmt.Println("==============================")
	fmt.Println("ID:           ", k.ID)
	fmt.Println("Fingerprint:  ", k.Fingerprint)
	fmt.Println("Name:         ", k.Name)
}
