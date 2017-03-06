package service

import (
	"encoding/json"
	"io/ioutil"

	"fmt"

	"errors"

	"github.com/mikaelm1/pirate/data"
)

// FetchDroplets retrieves all the droplets for the user (up to 100)
func (c *DOService) FetchDroplets(droplets *data.Droplets) error {
	res, err := c.MakeGETRequest("https://api.digitalocean.com/v2/droplets?page=1&per_page=100")
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &droplets)
	if err != nil {
		return err
	}
	// fmt.Println(os.Stdout, string(body))
	return nil
}

// FetchDroplet retrieves a single droplet. The droplet must have ID set
func (c *DOService) FetchDroplet(droplet *data.SingleDroplet, id int) error {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%v", id)
	res, err := c.MakeGETRequest(url)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &droplet)
	if err != nil {
		return err
	}
	return nil
}

// CreateDroplet creates a new droplet
func (c *DOService) CreateDroplet(droplet *data.SingleDroplet, dropletBody *data.DropletCreate) error {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets")
	body, err := json.Marshal(dropletBody)
	if err != nil {
		return err
	}
	// fmt.Println("Creating droplet with: ", droplet)
	res, err := c.MakePostRequest(url, body)
	if err != nil {
		return err
	}
	// fmt.Println(res.Header)
	body, err = ioutil.ReadAll(res.Body)
	// fmt.Println(os.Stdout, string(body))
	if err != nil {
		return err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &droplet)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDroplet will send a DELETE request to DO to delete a droplet
func (c *DOService) DeleteDroplet(id string) error {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%v", id)
	res, err := c.SendDeleteRequest(url)
	if err != nil {
		return err
	}
	if res.StatusCode != 204 {
		return errors.New("There was an error deleting the droplet. Make sure the droplet ID is correct")
	}
	return nil
}
