package slack

import (
	"errors"
	"time"
)

const (
	defaultHTTPTimeOut = 5 * time.Second
)

var (
	// ErrNoWebhookRegistered raised on no webhook url set on option
	ErrNoWebhookRegistered = errors.New("no webhook url set on option, request cancelled")
	// ErrEngineUsedWithoutNew raised on engine used without using New()
	ErrEngineUsedWithoutNew = errors.New("engine not configured yet, please use func New()")
)
