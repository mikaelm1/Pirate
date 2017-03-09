package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
