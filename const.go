package slack

import (
	"time"
)

const (
	defaultHTTPTimeOut = 5 * time.Second

	errNoWebhookRegistered  = "no webhook url set on option, request cancelled"
	errEngineUsedWithoutNew = "engine not configured yet, please use func New()"
)
