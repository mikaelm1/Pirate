package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mikaelm1/pirate/data"
)

// GetUserInfo gets info about the user
func (c *DOService) GetUserInfo(u *data.Account) error {
	res, err := c.MakeGETRequest("https://api.digitalocean.com/v2/account")
	if err != nil {
		return err
	}
	// var account data.Account
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	res.Body.Read(body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &u)
	if err != nil {
		// fmt.Println("Error: ", err)
		return err
	}
	return nil
}
