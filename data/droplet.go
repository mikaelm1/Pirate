package data

import (
	"fmt"
)

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
	SSHKey    []string `json:"ss_keys"`
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

type Kernel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type SingleDroplet struct {
	SDroplet Droplet `json:"droplet"`
}

type Droplets struct {
	DropletsList []Droplet `json:"droplets"`
	// Links        string    `json:"links"`
}

type Network struct {
	V4Networks `json:"v4"`
}

type V4Network struct {
	IPAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
	Type      string `json:"type"`
}

type V4Networks []V4Network

// PrintInfo displays info about the droplet
func (d *Droplet) PrintInfo() {
	fmt.Println("======================================")
	fmt.Println("Name:            ", d.Name)
	fmt.Println("ID:              ", d.ID)
	fmt.Println("Created:         ", d.CreatedAt)
	fmt.Println("Kernel:          ", d.DKernel.Name)
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
