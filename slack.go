package slack

import (
	"fmt"
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

		var resp string
		err := e.doJSON(http.MethodPost, url, pl, &resp)
		if err != nil {
			return err
		}
		if resp != "ok" {
			return fmt.Errorf("[go-slack] response is not ok : %s", resp)
		}
	}

	return nil

}
