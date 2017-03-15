package cmd

import (
	"fmt"

	"errors"

	"github.com/mikaelm1/pirate/data"
	"github.com/spf13/cobra"
)

var (
	balancerID            string
	balancerName          string
	balancerAlgo          string
	balancerRegion        string
	numForwardingRules    int
	entryProtocols        []string
	entryPorts            []int
	targetProtocols       []string
	targetPorts           []int
	certificateID         string
	tlsPassthrough        []bool
	healthProtocol        string
	healthPort            int
	healthPath            string
	healthCheckInterval   int
	healthThreshold       int
	healthResponseTimeout int
	unhealthyThreshold    int
	dropletIDs            []int
	stickyType            string
	stickyCookieName      string
	stickyTLS             string
)

// balancersCmd represents the balancers command
var balancersCmd = &cobra.Command{
	Use:   "balancers",
	Short: "Commands for load balancers",
	RunE:  fetchLoadBalancers,
}

var balancersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new load balancer",
	RunE:  createLoadBalancer,
}

var balancersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete load balancer",
	RunE:  deleteLoadBalancer,
}

var addDropletsCmd = &cobra.Command{
	Use:   "add-droplets",
	Short: "Add droplets to a load balancer",
	RunE:  addDroplets,
}

var removeDropletsCmd = &cobra.Command{
	Use:   "remove-droplets",
	Short: "Remove droplets from a load balancer",
	RunE:  removeDroplets,
}

func removeDroplets(*cobra.Command, []string) error {
	if balancerID == "" {
		return fmt.Errorf("Must provide balancer ID")
	}
	if len(dropletIDs) == 0 {
		return fmt.Errorf("Must provide at least one droplet id to add to the load balancer")
	}
	fmt.Println("Removing droplets from load balancer...")
	_, err := DOService.BalancerRemoveDroplets(balancerID, dropletIDs)
	if err != nil {
		return err
	}
	fmt.Printf("Removed %d droplets from your load balancer\n", len(dropletIDs))
	return nil
}

func addDroplets(*cobra.Command, []string) error {
	if balancerID == "" {
		return fmt.Errorf("Must provide balancer ID")
	}
	if len(dropletIDs) == 0 {
		return fmt.Errorf("Must provide at least one droplet id to add to the load balancer")
	}
	fmt.Println("Adding droplets to load balancer...")
	_, err := DOService.BalancerAddDroplets(balancerID, dropletIDs)
	if err != nil {
		return err
	}
	fmt.Printf("Added %d droplets to your load balancer\n", len(dropletIDs))
	return nil
}

func deleteLoadBalancer(*cobra.Command, []string) error {
	if balancerID == "" {
		return fmt.Errorf("Must provide valid load balancer ID")
	}
	_, err := DOService.DeleteLoadBalancer(balancerID)
	if err != nil {
		return err
	}
	fmt.Println("Deleted load balancer with id:", balancerID)
	return nil
}

func createLoadBalancer(*cobra.Command, []string) error {
	fmt.Println("Creating new load balancer...")
	err := validateCreationFlags()
	if err != nil {
		return err
	}
	fmt.Printf("Adding %d forwarding rules\n", numForwardingRules)
	var forwardingRules data.ForwardingRules
	for i := 0; i < numForwardingRules; i++ {
		rule := data.ForwardingRule{
			EntryProtocol:  entryProtocols[i],
			EntryPort:      entryPorts[i],
			TargetProtocol: targetProtocols[i],
			TargetPort:     targetPorts[i],
			CertificateID:  certificateID,
			TLSPassthrough: tlsPassthrough[i],
		}
		err := rule.IsValid()
		if err != nil {
			return fmt.Errorf("Invalid parameters passed into creating forwarding rule #%d\n%s", i, err.Error())
		}
		forwardingRules = append(forwardingRules, rule)
	}
	health := data.HealthCheck{
		Protocol:               healthProtocol,
		Port:                   healthPort,
		Path:                   healthPath,
		CheckIntervalSeconds:   healthCheckInterval,
		ResponseTimeoutSeconds: healthResponseTimeout,
		HealthyThreshold:       healthThreshold,
		UnhealthyThreshold:     unhealthyThreshold,
	}
	balancer := data.LoadBalancerCreate{
		Name:        balancerName,
		Algorithm:   balancerAlgo,
		Rules:       forwardingRules,
		HealthCheck: health,
		DropletIDs:  dropletIDs,
	}
	if len(dropletIDs) == 0 {
		balancer.Region = balancerRegion
	}
	sticky := data.StickySessions{
		Type:             stickyType,
		CookieName:       stickyCookieName,
		CookieTLSSeconds: stickyTLS,
	}
	err = sticky.IsValid()
	if err != nil {
		return err
	}
	balancer.StickySessions = sticky
	_, err = DOService.CreateLoadBalancer(&balancer)
	if err != nil {
		return err
	}
	balancer.PrintInfo()
	fmt.Println("New load balancer created. May take a few minutes until it's fully online")
	return nil
}

func validateCreationFlags() error {
	if balancerName == "" {
		return errors.New("Must provide a name for the load balancer")
	}
	if !isAlgoValid() {
		return errors.New("Load balancer algorithm must be either 'round_robin' or 'least_connnections'")
	}
	if !isRegionValid(balancerRegion) {
		return fmt.Errorf("(%s) is not a valid region", balancerRegion)
	}
	if numForwardingRules < 1 {
		return fmt.Errorf("Must provide at least one forwarding rule")
	}
	if len(entryProtocols) != numForwardingRules || len(entryPorts) != numForwardingRules || len(targetProtocols) != numForwardingRules || len(targetPorts) != numForwardingRules || len(tlsPassthrough) != numForwardingRules {
		return fmt.Errorf("The variable arrays for the forwarding rules must the match the number of rules being created")
	}
	return nil
}

func fetchLoadBalancers(*cobra.Command, []string) error {
	fmt.Println("Fetching your load balancers...")
	var balancers data.LoadBalancers
	_, err := DOService.FetchAllLoadBalancers(&balancers)
	if err != nil {
		return err
	}
	balancers.PrintInfo()
	return nil
}

func isAlgoValid() bool {
	if balancerAlgo == "round_robin" || balancerAlgo == "least_connnections" {
		return true
	}
	return false
}

func init() {
	RootCmd.AddCommand(balancersCmd)
	balancersCmd.AddCommand(balancersCreateCmd)
	balancersCmd.AddCommand(balancersDeleteCmd)
	balancersCmd.AddCommand(addDropletsCmd)
	balancersCmd.AddCommand(removeDropletsCmd)

	// create flags
	balancersCreateCmd.Flags().StringVarP(&balancerName, "name", "n", "", "Name of new load balancer")
	balancersCreateCmd.Flags().StringVarP(&balancerAlgo, "Algorithm", "a", "round_robin", "Algorithm for determining which backend droplet gets the client's request")
	balancersCreateCmd.Flags().StringVarP(&balancerRegion, "region", "r", "nyc3", "Region for load balancer")
	balancersCreateCmd.Flags().IntVar(&numForwardingRules, "num-rules", 1, "The number of forwarding rules to apply to the new load balancer")
	balancersCreateCmd.Flags().StringSliceVar(&entryProtocols, "entry-protocols", []string{"http"}, "Array of protocols for traffic to load balancer per forwarding rule being added")
	balancersCreateCmd.Flags().IntSliceVar(&entryPorts, "entry-ports", []int{80}, "Array of ports for traffic to load balancer per forwarding rule being added")
	balancersCreateCmd.Flags().StringSliceVar(&targetProtocols, "target-protocols", []string{"http"}, "Array of protocols for traffic from load balancer to droplets. One per forwarding rule")
	balancersCreateCmd.Flags().IntSliceVar(&targetPorts, "target-ports", []int{80}, "Array of ports for traffic from load balancer to droplets. One per forwarding rule")
	balancersCreateCmd.Flags().StringVarP(&certificateID, "cert-id", "c", "", "ID of TLS certificate for SSL termination")
	balancersCreateCmd.Flags().BoolSliceVarP(&tlsPassthrough, "tls-passthroughs", "t", []bool{false}, "Pass SSL encrypted traffic to droplets. One per forwarding rule")
	balancersCreateCmd.Flags().StringVar(&healthProtocol, "health-protocol", "http", "The protocol used for health checking droplets")
	balancersCreateCmd.Flags().IntVar(&healthPort, "health-port", 80, "The port for health checking droplets")
	balancersCreateCmd.Flags().StringVar(&healthPath, "health-path", "/", "The path on droplets listening for health checks")
	balancersCreateCmd.Flags().IntVar(&healthCheckInterval, "health-interval", 10, "Number of seconds between consecutive health checks")
	balancersCreateCmd.Flags().IntVar(&healthThreshold, "health-threshold", 5, "Number of healthy responses for droplet to be considered healthy")
	balancersCreateCmd.Flags().IntVar(&healthResponseTimeout, "health-response", 5, "Seconds load balancer will wait for response before marking health check as failing")
	balancersCreateCmd.Flags().IntVar(&unhealthyThreshold, "unhealthy", 3, "Number of failing checks before droplet is removed from pool")
	balancersCreateCmd.Flags().IntSliceVar(&dropletIDs, "droplet-ids", []int{}, "Array of droplet IDs to be assigned to load balancer")
	balancersCreateCmd.Flags().StringVar(&stickyType, "sticky-type", "none", "Should client requests be served by same droplet")
	balancersCreateCmd.Flags().StringVar(&stickyCookieName, "sticky-name", "", "If using sticky sessions, name to use for cookies")
	balancersCreateCmd.Flags().StringVar(&stickyTLS, "sticky-tls", "", "If using sticky sessions, number of seconds till cookie expires")

	// delete flags
	balancersDeleteCmd.Flags().StringVar(&balancerID, "balancer-id", "", "The ID of the load balancer to delete")

	// add droplets flags
	addDropletsCmd.Flags().StringVar(&balancerID, "balancer-id", "", "The ID of the load balancer to delete")
	addDropletsCmd.Flags().IntSliceVar(&dropletIDs, "droplet-ids", []int{}, "Array of droplet IDs to be assigned to load balancer")

	// remove droplets flags
	removeDropletsCmd.Flags().StringVar(&balancerID, "balancer-id", "", "The ID of the load balancer to delete")
	removeDropletsCmd.Flags().IntSliceVar(&dropletIDs, "droplet-ids", []int{}, "Array of droplet IDs to be assigned to load balancer")
}
