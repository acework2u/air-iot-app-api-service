package events

type Notification interface {
	Send() error
}

type notificationEvent struct {
	Topic string
	value string
}

type StartCleanOzoneEvent struct {
}

type EndCleanOzoneEvent struct {
}
type CancelClenOzoneEvent struct {
}

func NewNotificationEvent() Notification {
	return &notificationEvent{}
}

func (e *notificationEvent) Send() error {
	panic("No action")
}
