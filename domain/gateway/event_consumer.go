package gateway

import "context"

type ConsumeCallback func(key, value []byte) error

// EventConsumer consumes all kinds of events
type EventConsumer interface {
	ConsumeEvents(ctx context.Context, callback ConsumeCallback)
	Stop() error
}
