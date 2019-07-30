package ssdpdiscovery

import (
	"github.com/FantasticFiasco/axis-discovery-go/ssdpdiscovery"
)

type message struct {
	data	string
}

func newMessage(b []byte) *message {
	return &message{
		data: string(b),
	}
}

func (m *message) parseNotify() (d ssdpdiscovery.Device) {
	return
}
