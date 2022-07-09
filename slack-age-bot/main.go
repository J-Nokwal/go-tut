package main

import (
	_ "context"
	"fmt"
	_ "log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3400752009570-3394110455782-4qOcU0GEzU8S3V3HXTLbAhzG")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03BSN5H0HY-3400754387522-5842d0f01bf248c6bcd9d0d79c8ba0b595ebec805768a54c51d4889596ee64c8")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
}
