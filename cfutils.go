package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type SimpleDate struct {
	Year  int
	Month int
	Day   int
}

func NewCfRequest(endpoint string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s.campfirenow.com%s", Config.Domain, endpoint), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(Config.ApiKey, "X")
	req.Header.Add("User-Agent", "campfire-tools (https://github.com/vpetrenko/campfire-tools)")
	return req, nil
}

func PPJSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func ParseInt(data string) int {
	res, err := strconv.ParseInt(strings.TrimSpace(data), 10, 32)
	CheckFatalError(err, "Could not parse int")
	return int(res)
}

func ParseInt64(data string) int64 {
	res, err := strconv.ParseInt(strings.TrimSpace(data), 10, 64)
	CheckFatalError(err, "Could not parse int64")
	return res
}

func ParseCfTime(data string) time.Time {
	const FORMAT_STRING = "2006/01/02 15:04:05 -0700"

	dateTime, err := time.Parse(FORMAT_STRING, data)
	CheckFatalError(err, "Could not parse time")
	return dateTime
}

func ParseSlackTime(data string) time.Time {
	ts := strings.Split(data, ".")
	return time.Unix(ParseInt64(ts[0]), ParseInt64(ts[1]))
}

func ParseSimpleDate(data string) SimpleDate {
	const FORMAT_STRING = "2006-01-02"

	dateTime, err := time.Parse(FORMAT_STRING, data)
	CheckFatalError(err, "Could not parse simple date")
	return SimpleDate{dateTime.Year(), int(dateTime.Month()), dateTime.Day()}
}

func SimpleDateToTime(date SimpleDate) time.Time {
	const FORMAT_STRING = "2006/01/02"

	dateTime, err := time.Parse(FORMAT_STRING, fmt.Sprintf("%d/%02d/%02d", date.Year, date.Month, date.Day))
	CheckFatalError(err, "Could not parse SimpleDate")
	return dateTime
}
