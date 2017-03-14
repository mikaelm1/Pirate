package service

import (
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
