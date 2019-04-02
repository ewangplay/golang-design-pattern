package bridge

import "testing"

func TestSMSMessager(t *testing.T) {
	m := NewConcreteMessager(ViaSMS())
	m.SendMessage("have a drink?", "Bob")
}

func TestEmailMessager(t *testing.T) {
	m := NewConcreteMessager(ViaEmail())
	m.SendMessage("have a drink?", "Tom")
}
