package service

import (
	"net/http"

	"io/ioutil"

	"encoding/json"

	"github.com/mikaelm1/pirate/data"
)

// FetchAllImages sends a request to get all images
func (c *DOService) FetchAllImages(i *data.Images) (*http.Response, error) {
	res, err := c.MakeGETRequest("https://api.digitalocean.com/v2/images?page=1&per_page=50")
	if err != nil {
		return res, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &i)
	if err != nil {
		return res, err
	}
	return res, nil
}

// FetchAllDistroImages sends a request to get all distribution type images
func (c *DOService) FetchAllDistroImages(i *data.Images) (*http.Response, error) {
	res, err := c.MakeGETRequest("https://api.digitalocean.com/v2/images?type=distribution&page=1&per_page=50")
	if err != nil {
		return res, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &i)
	if err != nil {
		return res, err
	}
	return res, nil
}
