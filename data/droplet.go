package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Droplet is the model for a droplet
type Droplet struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Memory    int      `json:"memory"`
	Disk      int      `json:"disk"`
	Locked    bool     `json:"locked"`
	Status    string   `json:"status"`
	DKernel   Kernel   `json:"kernel"`
	CreatedAt string   `json:"created_at"`
	DNetwork  Network  `json:"networks"`
	Image     Image    `json:"image"`
	Features  []string `json:"features"`
}

// DropletCreate to be used when creating a new droplet
type DropletCreate struct {
	Name              string   `json:"name"`
	Region            string   `json:"region"`
	Size              string   `json:"size"`
	Image             string   `json:"image"`
	SSHKeys           []string `json:"ssh_keys"`
	IPV6              bool     `json:"ipv6"`
	PrivateNetworking bool     `json:"private_networking"`
	Backups           bool     `json:"backups"`
}

// Kernel is base model
type Kernel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// SingleDroplet is model for a single droplet
type SingleDroplet struct {
	SDroplet Droplet `json:"droplet"`
}

// Droplets is model for array of droplets
type Droplets struct {
	DropletsList []Droplet `json:"droplets"`
	// Links        string    `json:"links"`
}

// Network is base model for network
type Network struct {
	V4Networks `json:"v4"`
	V6Networks `json:"v6"`
}

// V4Network is base model for ip4 network
type V4Network struct {
	IPAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
	Type      string `json:"type"`
}

// V4Networks is array of ip4 networks
type V4Networks []V4Network

// V6Network is the base model ip6 network
type V6Network struct {
	IPAddress string `json:"ip_address"`
	Netmask   int    `json:"netmask"`
	Gateway   string `json:"gateway"`
	Type      string `json:"type"`
}

// V6Networks is model for an array of ip6 networks
type V6Networks []V6Network

// PrintInfo displays info about the droplet
func (d *Droplet) PrintInfo() {
	if viper.GetString("output") == "json" {
		d.JSONPrint()
	} else {
		d.TextPrint()
	}
}

// JSONPrint displays droplet info as JSON
func (d *Droplet) JSONPrint() {
	output, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println("Error parsing to JSON")
	}
	os.Stdout.Write(output)
}

// TextPrint displays droplet info as text
func (d *Droplet) TextPrint() {
	fmt.Println("======================================")
	fmt.Println("Name:            ", d.Name)
	fmt.Println("ID:              ", d.ID)
	fmt.Println("Created:         ", d.CreatedAt)
	fmt.Println("Image Distro:    ", d.Image.Distro)
	if len(d.DNetwork.V4Networks) > 0 {
		fmt.Println("IP4:             ", d.DNetwork.V4Networks[0].IPAddress)
	}
	// fmt.Println("V4 IP:   ", d.DNetwork.V4Networks[0].IPAddress)
}

func (d Droplets) Len() int {
	return len(d.DropletsList)
}

func (d Droplets) Less(i, j int) bool {
	return d.DropletsList[i].Name < d.DropletsList[j].Name
}

func (d Droplets) Swap(i, j int) {
	d.DropletsList[i], d.DropletsList[j] = d.DropletsList[j], d.DropletsList[i]
}
