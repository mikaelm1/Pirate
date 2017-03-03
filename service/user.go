package service

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mikaelm1/pirate/data"
)

// GetUserInfo gets info about the user
func (c *DOService) GetUserInfo() error {
	res, err := c.MakeGETRequest("https://api.digitalocean.com/v2/account")
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
