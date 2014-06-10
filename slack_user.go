package main

import (
	"encoding/json"
	"io/ioutil"
)

type SlackProfile struct {
	Email string
}

type SlackUser struct {
	Id      string
	Name    string
	Deleted bool
	Profile SlackProfile
}

func GetSlackUsers() []SlackUser {
	usersData, err := ioutil.ReadFile(Config.SlackData + "/users.json")
	CheckFatalError(err, "Could not read slack users")

	var users []SlackUser
	err = json.Unmarshal(usersData, &users)
	CheckFatalError(err, "Could not parse slack users file")
	return users
}
