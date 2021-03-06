package main

import (
	"context"
	"log"

	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient("<YOUR SLACK BOT TOKEN>")

	bot.Command("echo <word>", "Echo a word!", func(request slacker.Request, response slacker.ResponseWriter) {
		word := request.Param("word")

		attachments := []slack.Attachment{}
		attachments = append(attachments, slack.Attachment{
			Color:      "red",
			AuthorName: "Raed Shomali",
			Title:      "Attachment Title",
			Text:       "Attachment Text",
		})

		response.Reply(word, slacker.WithAttachments(attachments))
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
