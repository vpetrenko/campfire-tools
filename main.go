package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var flagFrom string
var flagTo string
var flagCfRoom string
var flagSlackRoom string

func init() {
	flag.StringVar(&flagFrom, "from", "2000-01-01", "-from YYYY-MM-DD")
	flag.StringVar(&flagTo, "to", time.Now().Format("2006-01-02"), "-to YYYY-MM-DD")
	flag.StringVar(&flagCfRoom, "cfroom", "", "")
	flag.StringVar(&flagSlackRoom, "slroom", "", "")
}

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " -cfroom CfRoomName -slroom SlackRoomName [-from YYYY-MM-DD] [-to YYYY-MM-DD] <import | check>")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		switch args[0] {
		case "import":
			fmt.Println("Running campfire crawler...")
			RunCrawler()
		case "check":
			if flagCfRoom == "" || flagSlackRoom == "" {
				printUsage()
				os.Exit(1)
			}
			fmt.Printf("\nCampfire Room: %s", flagCfRoom)
			fmt.Printf("\nSlack Room: %s", flagSlackRoom)
			fmt.Printf("\nFrom: %s", flagFrom)
			fmt.Printf("\nTo: %s", flagTo)
			fmt.Println("")
			fmt.Println("Checking messages...")
			CheckMessages(flagCfRoom, flagSlackRoom, ParseSimpleDate(flagFrom), ParseSimpleDate(flagTo))
		default:
			printUsage()
		}

	} else {
		printUsage()
	}
}
