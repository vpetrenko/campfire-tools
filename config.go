package main

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
)

type Configuration struct {
	DbName    string
	ApiKey    string
	Domain    string
	SlackData string
}

var Config Configuration

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err := decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Could not read config file", err)
	}
	refValue := reflect.ValueOf(&Config).Elem()
	typeOfT := refValue.Type()
	for i := 0; i < refValue.NumField(); i++ {
		f := refValue.Field(i)
		if f.String() == "" {
			log.Fatalf("%s could not be empty", typeOfT.Field(i).Name)
		}
	}
}

func CheckFatalError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
