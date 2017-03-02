package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mikaelm1/pirate/data"
)

// GetUserInfo gets info about the user
func (c *DOService) GetUserInfo() error {
	token := getUserToken()
	bearer := fmt.Sprintf("bearer %v", token)
	c.Client()
	req, err := http.NewRequest("GET", "https://api.digitalocean.com/v2/account", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", bearer)
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	var account data.Account
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	res.Body.Read(body)
	err = json.Unmarshal(body, &account)
	if err != nil {
		// fmt.Println("Error: ", err)
		return err
	}
	account.PrintInfo()
	return nil
}
