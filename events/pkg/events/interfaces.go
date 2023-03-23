package events

import "time"

type Event interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() any
}

type EventHandler interface {
	Handle(event Event)
}

type EventDispatcher interface {
	Register(eventName string, handler EventHandler) error
	Dispatch(event Event) error
	Unregister(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear() error
}
