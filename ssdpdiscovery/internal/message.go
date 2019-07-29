package ssdpdiscovery

import (
	"github.com/FantasticFiasco/axis-discovery-go/ssdpdiscovery"
	"net"
)

type message struct {
	addr	net.UDPAddr
	b		[]byte
}

func (m message) parse() (device ssdpdiscovery.Device) {
	return
}
