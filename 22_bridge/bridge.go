package bridge

import "log"

type Messager interface {
	SendMessage(text, to string)
}

type MessageChannel interface {
	Send(text, to string)
}

type SMS struct{}

func ViaSMS() MessageChannel {
	return &SMS{}
}

func (*SMS) Send(text, to string) {
	log.Printf("SMS to %v: %v", to, text)
}

type Email struct{}

func ViaEmail() MessageChannel {
	return &Email{}
}

func (*Email) Send(text, to string) {
	log.Printf("Email to %v: %v", to, text)
}

type ConcreteMessager struct {
	channel MessageChannel
}

func NewConcreteMessager(channel MessageChannel) *ConcreteMessager {
	return &ConcreteMessager{
		channel: channel,
	}
}

func (m *ConcreteMessager) SendMessage(text, to string) {
	m.channel.Send(text, to)
}
