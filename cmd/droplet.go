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
	//create droplet vairables
	dropletName    string
	regionName     string
	dropletSize    string
	imageName      string
	backupsEnabled bool
	sshKey         []string
	// delete droplet
	dropletID string
)

// dropletCmd represents the droplet command
var dropletCmd = &cobra.Command{
	Use:   "droplet",
	Short: "Run actions related to droplets.",
	RunE:  handleCommand,
}

var dropletCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create droplet",
	RunE:  handleCreate,
}

var dropletDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a droplet",
	RunE:  handleDelete,
}

func isImageNameValid(name string) bool {
	valids := [...]string{"debian-8-x64", "fedora-24-x64", "centos-6-x64", "fedora-25-x64", "ubuntu-16-04-x64"}
	for _, img := range valids {
		if img == name {
			return true
		}
	}
	return false
}

func isRegionValid(region string) bool {
	valids := [...]string{"nyc1", "nyc2", "nyc3", "sfo1", "sfo2", "ams2", "ams3", "fra1", "tor1", "sgp1", "lon1", "blr1"}
	for _, r := range valids {
		if region == r {
			return true
		}
	}
	return false
}

func isDropletSizeValid(size string) bool {
	valids := [...]string{"512mb", "1gb", "2gb", "4gb", "8gb", "16gb", "32gb", "48gb", "64gb"}
	for _, s := range valids {
		if size == s {
			return true
		}
	}
	return false
}

func handleCreate(*cobra.Command, []string) error {
	if dropletName == "" {
		return errors.New("Must name droplet")
	}
	if !isImageNameValid(imageName) {
		return errors.New("The image name you provided is invalid")
	}
	if !isDropletSizeValid(dropletSize) {
		errorMsg := fmt.Sprintf("%v is not a valid size", dropletSize)
		return errors.New(errorMsg)
	}
	if !isRegionValid(regionName) {
		errorMsg := fmt.Sprintf("%v is not a valid region name", regionName)
		return errors.New(errorMsg)
	}
	var droplet = data.DropletCreate{
		Name:              dropletName,
		Size:              dropletSize,
		Image:             imageName,
		Region:            regionName,
		PrivateNetworking: true,
		SSHKeys:           sshKey,
		IPV6:              true,
		Backups:           backupsEnabled,
	}
	var singleDroplet data.SingleDroplet
	err := DOService.CreateDroplet(&singleDroplet, &droplet)
	if err != nil {
		return err
	}
	fmt.Println("\nHere's your new droplet...")
	singleDroplet.SDroplet.PrintInfo()
	return nil
}

func handleDelete(*cobra.Command, []string) error {
	if dropletID == "" {
		return errors.New("Need the id of the droplet")
	}
	err := DOService.DeleteDroplet(dropletID)
	if err != nil {
		return err
	}
	fmt.Println("Your droplet has been successfully deleted")
	return nil
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
	dropletCmd.AddCommand(dropletCreateCmd)
	dropletCmd.AddCommand(dropletDeleteCmd)

	dropletCmd.Flags().BoolVarP(&listDropletsFlag, "list", "l", false, "list my droplets")
	dropletCmd.Flags().IntVarP(&getSingleDroplet, "single", "s", -1, "get droplet by id")
	// create droplet flags
	dropletCreateCmd.Flags().StringVarP(&dropletName, "name", "n", "", "name new droplet")
	dropletCreateCmd.Flags().StringVarP(&regionName, "region", "r", "nyc3", "name the region")
	dropletCreateCmd.Flags().StringVarP(&dropletSize, "size", "s", "512mb", "size of cpu")
	dropletCreateCmd.Flags().StringVarP(&imageName, "image", "i", "ubuntu-16-04-x64", "distribution image type")
	dropletCreateCmd.Flags().StringArrayVarP(&sshKey, "key", "k", []string{}, "ID's of your ssh keys")
	dropletCreateCmd.Flags().BoolVarP(&backupsEnabled, "backups", "b", false, "enable backups")
	// delete droplet flags
	dropletDeleteCmd.Flags().StringVarP(&dropletID, "droplet-id", "i", "", "droplet id")
}
