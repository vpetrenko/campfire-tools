package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type SlackMessage struct {
	Type   string
	User   string
	Text   string
	Ts     string
	Room   string
	TsTime time.Time
}

func GetSlackMessages(room string, year int, month int, day int) []SlackMessage {
	filePath := fmt.Sprintf("%s/%s/%d-%02d-%02d.json", Config.SlackData, room, year, month, day)
	messages := getSlackMessages(filePath)
	for i, m := range messages {
		messages[i].Room = room
		messages[i].TsTime = ParseSlackTime(m.Ts)
	}
	return messages
}

func getSlackMessages(filePath string) []SlackMessage {
	messagesData, err := ioutil.ReadFile(filePath)
	CheckFatalError(err, "Could not read slack messages")

	var messages []SlackMessage
	err = json.Unmarshal(messagesData, &messages)
	CheckFatalError(err, "Could not parse slack messages file "+filePath)
	return messages
}

func GetSlackMessagesForRoom(room string) []SlackMessage {
	msgFiles, err := ioutil.ReadDir(Config.SlackData + "/" + room)
	CheckFatalError(err, "Could not get message files")

	messages := make([]SlackMessage, 0)
	for _, mf := range msgFiles {
		if strings.HasSuffix(mf.Name(), ".json") {
			filePath := fmt.Sprintf("%s/%s/%s", Config.SlackData, room, mf.Name())
			dayMsgs := getSlackMessages(filePath)
			messages = append(messages, dayMsgs...)
		}
	}

	for i, _ := range messages {
		messages[i].Room = room
	}
	return messages
}

func GetAllSlackMessages() []SlackMessage {
	roomDirs, err := ioutil.ReadDir(Config.SlackData)
	CheckFatalError(err, "Could not get rooms names")
	messages := make([]SlackMessage, 0)
	for _, rd := range roomDirs {
		if rd.IsDir() {
			roomName := rd.Name()
			roomMsgs := GetSlackMessagesForRoom(roomName)
			messages = append(messages, roomMsgs...)
		}
	}
	return messages
}
