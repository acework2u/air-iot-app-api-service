package event

type Notification interface {
	SendMessage()
}

type notification struct {
	Topic   string
	Message string
}
