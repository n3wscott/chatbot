module github.com/n3wscott/chatbot

go 1.13

require (
	github.com/cloudevents/sdk-go v1.0.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/ko v0.3.0 // indirect
	github.com/n3wscott/prbuilder v0.0.1
	github.com/n3wscott/tomles v0.0.3
	github.com/nlopes/slack v0.6.0
)

replace github.com/nlopes/slack => github.com/n3wscott/nlopes-slack v0.2.1-0.20200217182150-f8647e88de75
