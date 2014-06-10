package main

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/ziutek/mymysql/godrv"
)

type Room struct {
	Id              int64
	Name            string
	Topic           string
	MembershipLimit int `json:"Membership_Limit"`
	Full            bool
	OpenToGuests    bool   `json:"Open_To_Guests"`
	UpdatedAt       int64  `json:"-"`
	Updated_At      string `db:"-"`
	CreatedAt       int64  `json:"-"`
	Created_At      string `db:"-"`
}

type User struct {
	Id           int64
	Name         string
	EmailAddress string `json:"Email_Address"`
	Admin        bool
	CreatedAt    string `json:"Created_At"`
	Type         string
	AvatarUrl    string `json:"Avatar_Url"`
}

type Message struct {
	Id         int64
	RoomId     int64 `json:"Room_Id"`
	UserId     int64 `json:"User_Id"`
	Body       string
	CreatedAt  int64  `json:"-"`
	Created_At string `db:"-"`
	Type       string
	Starred    bool
}

func InitDb() *gorp.DbMap {
	dbConnStr := fmt.Sprintf("tcp:localhost:3306*%s/root/", Config.DbName)
	db, err := sql.Open("mymysql", dbConnStr)
	CheckFatalError(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbmap.AddTableWithName(Room{}, "room").SetKeys(false, "Id")
	dbmap.AddTableWithName(Message{}, "message").SetKeys(false, "Id")
	dbmap.AddTableWithName(User{}, "user").SetKeys(false, "Id")

	err = dbmap.CreateTablesIfNotExists()
	CheckFatalError(err, "Create tables failed")

	return dbmap
}
