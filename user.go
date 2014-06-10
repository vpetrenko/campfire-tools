package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type userWrapper struct {
	User User
}

func GetUser(id int64) User {
	var endpoint string
	if id == 0 {
		endpoint = "/users/me.json"
	} else {
		endpoint = fmt.Sprintf("/users/%d.json", id)
	}
	req, err := NewCfRequest(endpoint)
	CheckFatalError(err, "Could not create request")

	resp, err := http.DefaultClient.Do(req)
	CheckFatalError(err, "Could not execute campfire request")

	body, err := ioutil.ReadAll(resp.Body)
	CheckFatalError(err, "Could not read campfire response")

	var user userWrapper
	err = json.Unmarshal(body, &user)
	CheckFatalError(err, "Could not parse campfire response")
	return user.User
}
