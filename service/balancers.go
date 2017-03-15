package service

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/mikaelm1/pirate/data"
)

// FetchAllLoadBalancers sends request to get all load balancers
func (c *DOService) FetchAllLoadBalancers(balancers *data.LoadBalancers) (*http.Response, error) {
	res, err := c.MakeGETRequest("https://api.digitalocean.com/v2/load_balancers")
	if err != nil {
		return res, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &balancers)
	if err != nil {
		return res, err
	}
	// fmt.Println(string(body))
	return res, nil
}

// CreateLoadBalancer sends post request to create a new load balancer
func (c *DOService) CreateLoadBalancer(balancer *data.LoadBalancer) (*http.Response, error) {
	body, err := json.Marshal(balancer)
	if err != nil {
		return nil, err
	}
	res, err := c.MakePostRequest("https://api.digitalocean.com/v2/load_balancers", body)
	if err != nil {
		return res, err
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Read(body)
	// fmt.Println(string(body))
	if res.StatusCode == 422 {
		return res, fmt.Errorf("Status Code 422: Make sure all protocols are on the same network layer")
	}
	if res.StatusCode >= 300 {
		return res, fmt.Errorf("Status Code %d: There was an error creating the load balancer", res.StatusCode)
	}
	balancers := data.LoadBalancers{
		Balancers: []data.LoadBalancer{*balancer},
	}
	err = json.Unmarshal(body, &balancers)
	if err != nil {
		return res, err
	}
	return res, nil
}

// DeleteLoadBalancer sends a request to delete a single load balancer
func (c *DOService) DeleteLoadBalancer(id string) (*http.Response, error) {
	res, err := c.SendDeleteRequest(fmt.Sprintf("https://api.digitalocean.com/v2/load_balancers/%s", id))
	if err != nil {
		return res, err
	}
	if res.StatusCode != 204 {
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Read(body)
		fmt.Println(string(body))
		return res, fmt.Errorf("Status Code %d: There was an error deleting load balancer with id: %s", res.StatusCode, id)
	}
	return res, nil
}
