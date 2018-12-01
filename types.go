package slack

import "net/http"

// Engine ..
type Engine struct {
	opt    Option
	client *http.Client
}

// Option ..
type Option struct {
	WebHookURL    string
	Channel       string
	CustomMessage string
}

type payload struct {
	Text string `json:"text"`
}
