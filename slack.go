package slack

import (
	"net/http"
)

// New function return configured go-slack engine.
// that already configured with default 5 seconds http timeout
func New(opt Option) *Engine {
	var (
		client *http.Client
	)

	client = &http.Client{
		Timeout: 5 * defaultHTTPTimeOut,
	}

	return &Engine{
		opt:    opt,
		client: client,
	}

}

// Send function post given message to webhook urls
func (e *Engine) Send(message string) error {

	if len(e.opt.WebHookURLs) < 1 {
		return ErrNoWebhookRegistered
	}

	if e.client == nil {
		return ErrEngineUsedWithoutNew
	}

	var (
		pl payload
	)

	pl.Text = message

	for _, url := range e.opt.WebHookURLs {
		err := e.doJSON(http.MethodPost, url, pl, nil)
		if err != nil {
			return err
		}
	}

	return nil

}
