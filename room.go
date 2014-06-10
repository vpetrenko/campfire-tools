package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type roomsWrapper struct {
	Rooms []Room
}

func GetRooms() []Room {
	req, err := NewCfRequest("/rooms.json")
	CheckFatalError(err, "Could not create request")

	resp, err := http.DefaultClient.Do(req)
	CheckFatalError(err, "Could not execute campfire request")

	body, err := ioutil.ReadAll(resp.Body)
	CheckFatalError(err, "Could not read campfire response")

	//  fmt.Print(string(body))
	var rooms roomsWrapper
	err = json.Unmarshal(body, &rooms)
	CheckFatalError(err, "Could not parse campfire response")
	for i, v := range rooms.Rooms {
		rooms.Rooms[i].UpdatedAt = ParseCfTime(v.Updated_At).Unix()
		rooms.Rooms[i].CreatedAt = ParseCfTime(v.Created_At).Unix()
	}
	return rooms.Rooms
}
