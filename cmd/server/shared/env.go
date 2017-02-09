package shared

import "github.com/getsentry/raven-go"

var (
	// Config contains flags passed to the API
	Config *Flags
	// Raven is sentry client
	Raven *raven.Client
)
