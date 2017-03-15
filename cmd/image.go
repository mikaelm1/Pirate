package cmd

import (
	"fmt"

	"github.com/mikaelm1/pirate/data"
	"github.com/spf13/cobra"
)

var (
	listAllImages    bool
	listDistroImages bool
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Image commands",
	RunE:  handleImageCommand,
}

func handleImageCommand(*cobra.Command, []string) error {
	fmt.Println("Fetching images...")
	var images data.Images
	if listAllImages {
		_, err := DOService.FetchAllImages(&images)
		if err != nil {
			return nil
		}
		fmt.Printf("Showing %d images:\n", len(images.Images))
		images.PrintInfo()
	} else {
		return fmt.Errorf("Must provide either (-l) or (-d) flag")
	}
	return nil
}

func init() {
	RootCmd.AddCommand(imageCmd)

	imageCmd.Flags().BoolVarP(&listAllImages, "list", "l", false, "List all of your images")
	imageCmd.Flags().BoolVarP(&listDistroImages, "distro", "d", false, "List all distribution images")
}
