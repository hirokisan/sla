package main

import (
	"github.com/nlopes/slack"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app     = kingpin.New("chat", "A command-line chat application.")
	post    = kingpin.Command("post", "post message to slack")
	url     = post.Flag("url", "incomming webhook url").Short('u').String()
	message = post.Flag("message", "message").Short('m').String()
)

func main() {
	switch kingpin.Parse() {
	case "post":
		if *url != "" {
			message := Message(*message)
			slack.PostWebhook(*url, message)
		}
	}
}

func Message(m string) *slack.WebhookMessage {
	return &slack.WebhookMessage{
		Text:        m,
		Attachments: []slack.Attachment{},
	}
}
