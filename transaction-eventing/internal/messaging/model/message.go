package model

type Message struct {
	Topic       string
	EventType   string
	Message     interface{}
	ServiceName string
	LogLevel    string
}
