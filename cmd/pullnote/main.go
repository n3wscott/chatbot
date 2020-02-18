package main

import (
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/types"
	"github.com/google/go-github/github"
)

func main() {
	ctx := context.Background()

	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Printf("will listen on :8080\n")
	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, gotEvent))
}

func gotEvent(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) {

	diff := &GitHubDiff{}
	if err := event.DataAs(diff); err != nil {
		fmt.Println("Got a github event error", err.Error())
		return
	}

	action, err := types.ToString(event.Extensions()["action"])
	if err != nil {
		fmt.Println("Failed to get action", err.Error())
		return
	}
	fmt.Println("Got a github event of", action)

	switch action {
	case "new":
		fmt.Printf("GOT A NEW PR!\n")

		r := cloudevents.NewEvent("1.0")
		r.SetSource("pullnote")
		r.SetType("slackbot.response")
		r.SetDataContentType("application/json")
		_ = r.SetData(&Message{
			Channel: channelFor(diff.PR),
			Text:    fmt.Sprintf("Updated PR: %s", diff.PR.GetHTMLURL()),
		})

		resp.RespondWith(200, &r)

	case "diff":
		fmt.Printf("GOT A PR UPDATE!\n")

		r := cloudevents.NewEvent("1.0")
		r.SetSource("pullnote")
		r.SetType("slackbot.response")
		r.SetDataContentType("application/json")
		_ = r.SetData(&Message{
			Channel: channelFor(diff.PR),
			Text:    fmt.Sprintf("Updated PR: %s", diff.PR.GetHTMLURL()),
		})

		resp.RespondWith(200, &r)
	}

	fmt.Printf("%s", event)
	//fmt.Printf("%s\n", cloudevents.HTTPTransportContextFrom(ctx))
}

func channelFor(pr *github.PullRequest) string {
	return "DBF21M2KW" // TODO: yolo
}

type Message struct {
	Channel string `json:"channel,omitempty"`
	Text    string `json:"text,omitempty"`
}

type GitHubDiff struct {
	PR   *github.PullRequest `json:"pr"`
	Diff string              `json:"diff"`
}
