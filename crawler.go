package main

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/davecgh/go-spew/spew"
	"time"
)

func PackDate(date time.Time) int {
	return date.Day() + int(date.Month())<<5 + date.Year()<<9
}

func CrawlMessages(dbmap *gorp.DbMap, room Room) {
	timeFrom := time.Unix(room.CreatedAt, 0)
	timeTo := time.Now()

	fmt.Println("Time from: ", timeFrom)
	fmt.Println("Time to: ", timeTo)

	oneDay, err := time.ParseDuration("24h")
	CheckFatalError(err, "Could not get day duration")

	_, err = dbmap.Exec("delete from message where RoomId = ?", room.Id)
	CheckFatalError(err, "Could not empty message table")
	for ti := timeFrom; PackDate(ti) <= PackDate(timeTo); ti = ti.Add(oneDay) {
		messages := GetMessages(room.Id, ti.Year(), int(ti.Month()), ti.Day())
		for _, m := range messages {
			err := dbmap.Insert(&m)
			CheckFatalError(err, "Could not add message to DB "+fmt.Sprintf("%v", m))
		}
		fmt.Println(ti)
		//		time.Sleep(10 * time.Millisecond)
	}
}

func CollectUsers(dbmap *gorp.DbMap) {
	var messages []Message
	_, err := dbmap.Select(&messages, "select * from message group by UserId")
	CheckFatalError(err, "select user ids from messages")

	_, err = dbmap.Exec("delete from user")
	for _, m := range messages {
		if m.Id != 0 {
			user := GetUser(m.UserId)
			spew.Dump(user)
			dbmap.Insert(&user)
		}
	}
}

func RunCrawler() {
	dbmap := InitDb()
	defer dbmap.Db.Close()

	fmt.Println("Getting campfire rooms...")
	rooms := GetRooms()
	_, err := dbmap.Exec("delete from room")
	CheckFatalError(err, "Could not empty room table")
	for _, v := range rooms {
		err := dbmap.Insert(&v)
		CheckFatalError(err, "Could not store rooms to DB")
	}
	PPJSON(rooms)

	for _, v := range rooms {
		fmt.Printf("\nGetting messages for %s room...\n", v.Name)
		CrawlMessages(dbmap, v)
		time.Sleep(20 * time.Millisecond)
	}

	fmt.Println("Collecting user list from messages...")
	CollectUsers(dbmap)
}
