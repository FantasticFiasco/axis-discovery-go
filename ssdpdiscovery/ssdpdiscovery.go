// Package ssdpdiscovery provides means to discover devices from Axis Communications on the network using SSDP.
package ssdpdiscovery

import (
	"github.com/pkg/errors"
	"net"
)

var multicastAddress = net.UDPAddr {
	IP:   []byte{ 239, 255, 255, 250 },
	Port: 1900,
}

func ListenPassive(f func(d Device)) error {
	conn, err := net.ListenMulticastUDP("udp", nil, &multicastAddress)
	if err != nil {
		return errors.Wrap(err, "Failed to listen to multicast address")
	}
	defer conn.Close()
	b := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(b)
		if err != nil {
			return errors.Wrap(err, "Failed to read from UDP connection")
		}
		m := newMessage(addr, b[:n])
		d, err := m.parseNotify()
		if err == nil {
			f(d)
		}
	}
}
