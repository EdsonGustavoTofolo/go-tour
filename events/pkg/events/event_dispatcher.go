package events

import (
	"errors"
)

var (
	ErrHandlerAlreadyRegistered = errors.New("Handler already registered")
)

type EventDispatcherImpl struct {
	handlers map[string][]EventHandler
}

func NewEventDispatcher() *EventDispatcherImpl {
	return &EventDispatcherImpl{
		handlers: make(map[string][]EventHandler),
	}
}

func (e *EventDispatcherImpl) Register(eventName string, handler EventHandler) error {
	if handlers, ok := e.handlers[eventName]; ok {
		for _, h := range handlers {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	e.handlers[eventName] = append(e.handlers[eventName], handler)

	return nil
}

func (e *EventDispatcherImpl) Dispatch(event Event) error {
	if handlers, ok := e.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	return nil
}

func (e *EventDispatcherImpl) Unregister(eventName string, handler EventHandler) error {
	if handlers, ok := e.handlers[eventName]; ok {
		for i, h := range handlers {
			if h == handler {
				e.handlers[eventName] = append(e.handlers[eventName][:i], e.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (e *EventDispatcherImpl) Has(eventName string, handler EventHandler) bool {
	if handlers, ok := e.handlers[eventName]; ok {
		for _, h := range handlers {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (e *EventDispatcherImpl) Clear() error {
	e.handlers = make(map[string][]EventHandler)
	return nil
}
