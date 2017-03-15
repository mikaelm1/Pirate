package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Image is the base model for an image
type Image struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Distro      string   `json:"distribution"`
	Slug        string   `json:"slug"`
	Public      bool     `json:"public"`
	Regions     []string `json:"regions"`
	CreatedAt   string   `json:"created_at"`
	Type        string   `json:"snapshot"`
	MinDiskSize int      `json:"min_disk_size"`
	ImageSize   float64  `json:"size_gigabytes"`
}

// Images is the model for an array of Image objects
type Images struct {
	Images []Image `json:"images"`
}

// SingleImage is the model for a single Image object
type SingleImage struct {
	Image Image `json:"image"`
}

// PrintInfo displays info about an array of Image objects
func (i *Images) PrintInfo() {
	for k := 0; k < len(i.Images); k++ {
		i.Images[k].PrintInfo()
	}
}

// PrintInfo displays info about the image
func (i *Image) PrintInfo() {
	if viper.GetString("output") == "json" {
		i.JSONPrint()
	} else {
		i.TextPrint()
	}
}

// JSONPrint displays image info as JSON
func (i *Image) JSONPrint() {
	output, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		fmt.Println("Error parsing to JSON")
	}
	os.Stdout.Write(output)
}

// TextPrint displays image info as text
func (i *Image) TextPrint() {
	fmt.Println("============ Image =============")
	fmt.Printf("ID:                %d\n", i.ID)
	fmt.Printf("Name:              %s\n", i.Name)
	fmt.Printf("Slug:              %s\n", i.Slug)
	fmt.Printf("Distro:            %s\n", i.Distro)
}
