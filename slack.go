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

// Send function post given message to webhook url
func (e *Engine) Send(message string) error {

	if e.opt.WebHookURL == "" {
		return ErrNoWebhookRegistered
	}

	if e.client == nil {
		return ErrEngineUsedWithoutNew
	}

	var (
		err error
		pl  payload
	)

	pl.Text = message

	err = e.doJSON(http.MethodPost, e.opt.WebHookURL, pl, nil)

	return err

}
