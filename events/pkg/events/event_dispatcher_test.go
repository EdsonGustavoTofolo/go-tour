package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload any
}

func (t *TestEvent) GetName() string {
	return t.Name
}

func (t *TestEvent) GetPayload() any {
	return t.Payload
}

func (t *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event Event) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcherImpl
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.eventDispatcher = NewEventDispatcher()
	s.handler = TestEventHandler{1}
	s.handler2 = TestEventHandler{2}
	s.handler3 = TestEventHandler{3}
	s.event = TestEvent{"EventTest1", "body"}
	s.event2 = TestEvent{"EventTest2", "body"}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	assert.Equal(s.T(), &s.handler, s.eventDispatcher.handlers[s.event.GetName()][0])
	assert.Equal(s.T(), &s.handler2, s.eventDispatcher.handlers[s.event.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Equal(ErrHandlerAlreadyRegistered, err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event 1
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Event 2
	err = s.eventDispatcher.Register(s.event2.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event2.GetName()]))

	err = s.eventDispatcher.Clear()
	s.NoError(err)
	s.Equal(0, len(s.eventDispatcher.handlers))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// Event 1
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	assert.True(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handler))
	assert.True(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handler2))
	assert.False(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handler3))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Unregister() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Unregister(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
	assert.Equal(s.T(), &s.handler, s.eventDispatcher.handlers[s.event.GetName()][0])
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event Event) {
	m.Called(event)
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eventHandler := &MockHandler{}
	eventHandler.On("Handle", &s.event)

	err := s.eventDispatcher.Register(s.event.GetName(), eventHandler)
	s.NoError(err)

	err = s.eventDispatcher.Dispatch(&s.event)
	s.NoError(err)

	eventHandler.AssertExpectations(s.T())
	eventHandler.AssertNumberOfCalls(s.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
