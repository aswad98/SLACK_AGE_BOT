package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)

	}
}

func main() {
	os.Setenv("SLACK-BOT-TOKEN", "xoxb-3714417788930-3738520890896-gIsz4UaYt1DUjCKCc9KLsKMD")
	os.Setenv("SLACK-APP-TOKEN", "xapp-1-A03LU2Q9GUW-3700105659447-3def75ad0daa676cff9f5ed27c9d88d4e07d138816f10a8f03cae8774a281be2")
	bot := slacker.NewClient(os.Getenv("SLACK-BOT-TOKEN"), os.Getenv("SLACK-APP-TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 1998",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Fatal(err)
			}
			age := 2022 - yob
			result := fmt.Sprintf("age is %d", age)
			response.Reply(result)
		},
	})

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
