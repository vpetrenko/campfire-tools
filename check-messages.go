package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"strings"
	"time"
)

func CheckMessages(cfRoom string, slackRoom string, from SimpleDate, to SimpleDate) {
	dbmap := InitDb()
	defer dbmap.Db.Close()

	slackMessages := GetSlackMessagesForRoom(slackRoom)
	lookup := make(map[int64]SlackMessage)
	for _, v := range slackMessages {
		ts := strings.Split(v.Ts, ".")
		key := ParseInt64(ts[0])
		lookup[key] = v
	}

	var cfRooms []Room
	_, err := dbmap.Select(&cfRooms, "select * from room where Name = ?", cfRoom)
	CheckFatalError(err, "Could not get room "+cfRoom)
	if len(cfRooms) != 1 {
		log.Fatalln("Could not get data for room " + cfRoom)
	}

	oneDay, _ := time.ParseDuration("24h")
	fromDate := SimpleDateToTime(from)
	toDate := SimpleDateToTime(to).Add(oneDay)

	var cfMessages []Message
	_, err = dbmap.Select(&cfMessages, "select * from message where UserId <> 0 AND RoomId = ? AND CreatedAt BETWEEN ? AND ? order by CreatedAt",
		cfRooms[0].Id, fromDate.Unix(), toDate.Unix())
	CheckFatalError(err, "Could not get campfire messages")
	totalCfMessages := 0
	totalSlackMessages := 0
	matched := make([]string, 0)
	for _, m := range cfMessages {
		createdAt := time.Unix(m.CreatedAt, 0)
		if m.Type == "TextMessage" {
			totalCfMessages++
			sm, ok := lookup[createdAt.Unix()]
			if ok {
				matched = append(matched, sm.Ts)
				totalSlackMessages++
			} else {
				spew.Dump(m)
				spew.Dump(createdAt)
			}
		}
	}

	fmt.Println("CF Messages: ", totalCfMessages)
	fmt.Println("Slack Messages: ", totalSlackMessages)
	if totalCfMessages == totalSlackMessages {
		fmt.Println("All messages match.")
	} else {
		fmt.Println("Some messages are missing. See below.")
	}
}
