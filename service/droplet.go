package service

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mikaelm1/pirate/data"
)

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
		// fmt.Println("Error: ", err)
		return err
	}
	// fmt.Println(os.Stdout, string(body))
	return nil
}
