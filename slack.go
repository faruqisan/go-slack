package slack

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
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

	if err := e.validateClient(); err != nil {
		return err
	}

	for _, url := range e.opt.WebHookURLs {
		err := e.send(http.MethodPost, url, message)
		if err != nil {
			return err
		}
	}

	return nil

}

// SendAsync function will send to all registered webhooks asynchronously
// returning err chan for error info, channel will be closed when there's any error
// or all process is done
func (e *Engine) SendAsync(message string) chan error {

	var (
		errChan = make(chan error, 1)
		err     error
		wg      sync.WaitGroup
	)

	if err = e.validateClient(); err != nil {
		errChan <- err
		close(errChan)
		return errChan
	}

	for _, url := range e.opt.WebHookURLs {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			err := e.send(http.MethodPost, u, message)
			if err != nil {
				errChan <- err
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	return errChan
}

func (e *Engine) send(httpMethod, url string, message string) error {

	var (
		pl payload
	)

	pl.Text = e.opt.CustomMessage + message

	resp, err := e.doString(httpMethod, url, pl)
	if err != nil {
		return err
	}

	// remove whitespace
	resp = strings.ReplaceAll(resp, " ", "")

	if resp != "ok" {
		return fmt.Errorf("[go-slack] response is not ok : %s", resp)
	}

	return nil
}

func (e *Engine) validateClient() error {
	if len(e.opt.WebHookURLs) < 1 {
		return ErrNoWebhookRegistered
	}

	if e.client == nil {
		return ErrEngineUsedWithoutNew
	}

	return nil
}
