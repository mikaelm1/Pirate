package data

import (
	"encoding/json"
	"fmt"
	"os"

	"errors"

	"github.com/spf13/viper"
)

// LoadBalancer is the base model
type LoadBalancer struct {
	ID              string          `json:"id"`
	StringRegion    string          `json:"regions"`
	Region          Region          `json:"region"`
	Name            string          `json:"name"`
	IP              string          `json:"ip"`
	Algorithm       string          `json:"algorithm"`
	Status          string          `json:"new"`
	CreatedAt       string          `json:"created_at"`
	RedirectToHTTPS bool            `json:"redirect_to_https"`
	DropletIDs      []int           `json:"droplet_ids"`
	Rules           ForwardingRules `json:"forwarding_rules"`
	HealthCheck     HealthCheck     `json:"health_check"`
	StickySessions  StickySessions  `json:"sticky_sessions"`
}

// LoadBalancerCreate is the model to use when creating a load balancer
type LoadBalancerCreate struct {
	ID              string          `json:"id"`
	StringRegion    string          `json:"regions"`
	Name            string          `json:"name"`
	IP              string          `json:"ip"`
	Algorithm       string          `json:"algorithm"`
	Status          string          `json:"new"`
	CreatedAt       string          `json:"created_at"`
	RedirectToHTTPS bool            `json:"redirect_to_https"`
	DropletIDs      []int           `json:"droplet_ids"`
	Rules           ForwardingRules `json:"forwarding_rules"`
	HealthCheck     HealthCheck     `json:"health_check"`
	StickySessions  StickySessions  `json:"sticky_sessions"`
	Region          string          `json:"region"`
}

// LoadBalancers is the model for an array of LoadBalancer objects
type LoadBalancers struct {
	Balancers []LoadBalancer `json:"load_balancers"`
}

// LoadBalancersCreate is the model for an array of LoadBalancerCreate objects
type LoadBalancersCreate struct {
	Balancers []LoadBalancerCreate `json:"load_balancers"`
}

// Region model
type Region struct {
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	Sizes     []string `json:"sizes"`
	Features  []string `json:"features"`
	Available bool     `json:"available"`
}

// ForwardingRule model
type ForwardingRule struct {
	EntryProtocol  string `json:"entry_protocol"` // http, https, or tcp. Traffic to balancer
	EntryPort      int    `json:"entry_port"`
	TargetProtocol string `json:"target_protocol"` // http, https, or tcp. Traffic from balancer to server
	TargetPort     int    `json:"target_port"`
	CertificateID  string `json:"certificate_id"`
	TLSPassthrough bool   `json:"tls_passthrough"`
}

// ForwardingRules array of ForwardingRule
type ForwardingRules []ForwardingRule

// HealthCheck model
type HealthCheck struct {
	Protocol               string `json:"protocol"` // http or tcp
	Port                   int    `json:"port"`
	Path                   string `json:"path"`
	CheckIntervalSeconds   int    `json:"check_interval_seconds"`
	ResponseTimeoutSeconds int    `json:"response_timeout_seconds"`
	HealthyThreshold       int    `json:"healthy_threshold"`
	UnhealthyThreshold     int    `json:"unhealthy_threshold"`
}

// StickySessions model
type StickySessions struct {
	Type             string `json:"type"` // cookies or none(default)
	CookieName       string `json:"cookie_name"`
	CookieTLSSeconds string `json:"cookie_tls_seconds"`
}

// IsValid validates ForwardingRule fields
func (r *ForwardingRule) IsValid() error {
	if r.EntryProtocol != "http" && r.EntryProtocol != "https" && r.EntryProtocol != "tcp" {
		return fmt.Errorf("A forwarding rule's entry protocol must be either http, https, or tcp. (%s) is not a valid option", r.EntryProtocol)
	}
	if r.TargetProtocol != "http" && r.TargetProtocol != "https" && r.TargetProtocol != "tcp" {
		return fmt.Errorf("A forwarding rule's target protocol must be either http, https, or tcp. (%s) is not a valid protocol", r.TargetProtocol)
	}
	return nil
}

// IsValid validates StickySessions fields
func (s *StickySessions) IsValid() error {
	if s.Type == "" || s.Type == "none" {
		return nil
	}
	if s.CookieName == "" {
		return errors.New("Cookie name must be set if using cookies for sticky sessions")
	}
	if s.CookieTLSSeconds == "" {
		return errors.New("Seconds till cookie expires must be set when using cookies for sticky sessions")
	}
	return nil
}

// PrintInfo displays info about a load balancer
func (b *LoadBalancerCreate) PrintInfo() {
	if viper.GetString("output") == "json" {
		b.JSONPrint()
	} else {
		b.TextPrint()
	}
}

// JSONPrint displays info in JSON format
func (b *LoadBalancerCreate) JSONPrint() {
	output, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		fmt.Println("Error parsing to JSON")
	}
	os.Stdout.Write(output)
}

// TextPrint displays info in text format
func (b *LoadBalancerCreate) TextPrint() {
	fmt.Println("========Load Balancer==============")
	fmt.Printf("ID:            %s\n", b.ID)
	fmt.Printf("NAME           %s\n", b.Name)
	fmt.Printf("IP:            %s\n", b.IP)
	fmt.Printf("Algorithm      %s\n", b.Algorithm)
	fmt.Printf("Status:        %s\n", b.Status)
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
	fmt.Println("========Load Balancer==============")
	fmt.Printf("ID:            %s\n", b.ID)
	fmt.Printf("NAME           %s\n", b.Name)
	fmt.Printf("IP:            %s\n", b.IP)
	fmt.Printf("Algorithm      %s\n", b.Algorithm)
	fmt.Printf("Status:        %s\n", b.Status)
}
