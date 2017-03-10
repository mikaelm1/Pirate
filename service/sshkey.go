package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"

	"errors"

	"github.com/mikaelm1/pirate/data"
)

// FetchAllSSHKeys retrieves all the ssh keys for the user
func (c *DOService) FetchAllSSHKeys(keys *data.ArraySSHKey) (*http.Response, error) {
	res, err := c.MakeGETRequest("https://api.digitalocean.com/v2/account/keys?page=1&per_page=100")
	if err != nil {
		return res, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &keys)
	if err != nil {
		return res, err
	}
	// fmt.Println(os.Stdout, string(body))
	return res, nil
}

// FetchSingleKey retrieves a single ssh key given an id or fingerprint
func (c *DOService) FetchSingleKey(id string, key *data.SingleSSHKey) (*http.Response, error) {
	res, err := c.MakeGETRequest(fmt.Sprintf("https://api.digitalocean.com/v2/account/keys/%s", id))
	if err != nil {
		return res, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &key)
	if err != nil {
		return res, err
	}
	return res, nil
}

// CreateSSHKey sends the request to create an SSH key
func (c *DOService) CreateSSHKey(key *data.SSHKey, singleKey *data.SingleSSHKey) (*http.Response, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/account/keys")
	body, err := json.Marshal(key)
	if err != nil {
		return nil, err
	}
	res, err := c.MakePostRequest(url, body)
	if err != nil {
		return res, err
	}
	if res.StatusCode == 422 {
		return res, errors.New("That public key is already in use")
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return res, err
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &singleKey)
	if err != nil {
		return res, err
	}
	return res, nil
}
