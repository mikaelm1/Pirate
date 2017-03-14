package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// LoadBalancer is the base model
type LoadBalancer struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	IP              string `json:"ip"`
	Algorithm       string `json:"algorithm"`
	Status          string `json:"new"`
	CreatedAt       string `json:"created_at"`
	RedirectToHTTPS bool   `json:"redirect_to_https"`
}

// LoadBalancers is model for an array of LoadBalancer objects
type LoadBalancers struct {
	Balancers []LoadBalancer `json:"load_balancers"`
}

// PrintInfo displays array of load balancers
func (b *LoadBalancers) PrintInfo() {
	if len(b.Balancers) == 0 {
		fmt.Println("You don't have any load balancers")
	}
	for i := 0; i < len(b.Balancers); i++ {
		b.Balancers[i].PrintInfo()
	}
}

// PrintInfo displays info about a load balancer
func (b *LoadBalancer) PrintInfo() {
	fmt.Println("Printing")
	if viper.GetString("output") == "json" {
		b.JSONPrint()
	} else {
		b.TextPrint()
	}
}

// JSONPrint displays info in JSON format
func (b *LoadBalancer) JSONPrint() {
	output, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		fmt.Println("Error parsing to JSON")
	}
	os.Stdout.Write(output)
}

// TextPrint displays info in text format
func (b *LoadBalancer) TextPrint() {

}
