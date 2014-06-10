package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type messagesWrapper struct {
	Messages []Message
}

func GetMessages(roomId int64, year int, month int, day int) []Message {
	req, err := NewCfRequest(fmt.Sprintf("/room/%d/transcript/%d/%d/%d.json", roomId, year, month, day))
	CheckFatalError(err, "Could not create request")

	resp, err := http.DefaultClient.Do(req)
	CheckFatalError(err, "Could not execute campfire request")

	body, err := ioutil.ReadAll(resp.Body)
	CheckFatalError(err, "Could not read campfire response")

	var messages messagesWrapper
	err = json.Unmarshal(body, &messages)
	CheckFatalError(err, "Could not parse campfire response")

	for i, v := range messages.Messages {
		messages.Messages[i].CreatedAt = ParseCfTime(v.Created_At).Unix()
	}

	return messages.Messages
}
