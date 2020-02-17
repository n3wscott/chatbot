package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/nlopes/slack"
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
	b, ok := event.Data.([]byte)
	if !ok {
		fmt.Printf("failed to get data as []byte]: %T\n", event.Data)
		return
	}
	rawEvent := json.RawMessage(b)

	se, err := parseRawEvent(rawEvent)
	if err != nil {
		fmt.Printf("failed handle raw event: %s\n", err.Error())
	}

	switch s := se.(type) {
	case *slack.MessageEvent:
		fmt.Printf("GOT A SLACK MESSAGE!\n")

		r := cloudevents.NewEvent("1.0")
		r.SetSource("echo")
		r.SetType("slackbot.response")
		r.SetDataContentType("application/json")
		_ = r.SetData(&Message{
			Channel: s.Channel,
			Text:    s.Text,
		})

		resp.RespondWith(200, &r)
	}

	fmt.Printf("%s", event)
	//fmt.Printf("%s\n", cloudevents.HTTPTransportContextFrom(ctx))
}

type Message struct {
	Channel string `json:"channel,omitempty"`
	Text    string `json:"text,omitempty"`
}

func parseRawEvent(rawEvent json.RawMessage) (interface{}, error) {
	event := &slack.Event{}
	err := json.Unmarshal(rawEvent, event)
	if err != nil {
		//rtm.IncomingEvents <- RTMEvent{"unmarshalling_error", &UnmarshallingErrorEvent{err}}
		return nil, err
	}

	v, exists := slack.EventMapping[event.Type]
	if !exists {
		return nil, fmt.Errorf("RTM Error: Received unmapped event %q: %s", event.Type, string(rawEvent))
	}
	t := reflect.TypeOf(v)
	recvEvent := reflect.New(t).Interface()
	if err := json.Unmarshal(rawEvent, recvEvent); err != nil {
		return nil, fmt.Errorf("RTM Error: Could not unmarshall event %q[%s]: %s", event.Type, string(rawEvent), err)
	}

	fmt.Println("GOT AN EVENT:", recvEvent)

	return recvEvent, nil
}
