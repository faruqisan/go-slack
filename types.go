package slack

import "net/http"

// Engine struct act as function receiver
type Engine struct {
	opt    Option
	client *http.Client
}

// Option struct define configuration for engine
type Option struct {
	WebHookURLs   []string
	Channel       string
	CustomMessage string
}

type payload struct {
	Text string `json:"text"`
}
