package mq

// Helper sends events to the message queue.
type Helper interface {
	SendEvent(key string, value []byte) (bool, error)
}
