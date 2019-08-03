// Package ssdpdiscovery provides means to discover network attached devices from Axis Communications using SSDP.
package ssdpdiscovery

import (
	"github.com/pkg/errors"
	"net"
)

// The multicast address where SSDP notification are announced
var multicastAddr = net.UDPAddr{
	IP:   []byte{239, 255, 255, 250},
	Port: 1900,
}

type AliveHandler func(d *Device)

type ByeByeHandler func(d *Device)

// ListenPassive will passively listen for SSDP notifications on the network.
func ListenPassive(alive AliveHandler, byeBye ByeByeHandler) error {
	conn, err := net.ListenMulticastUDP("udp", nil, &multicastAddr)
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
		m := parseMessage(b[:n])
		if m.method == notify ||
			m.headers[nt] != "urn:axis-com:service:BasicService:1" {
			continue
		}
		if m.headers[nts] == "ssdp:alive" {
			alive(toDevice(addr, m))
		} else if m.headers[nts] == "ssdp:byebye" {
			byeBye(toDevice(addr, m))
		}
	}
}

func toDevice(addr *net.UDPAddr, m *message) *Device {
	return &Device{
		//Addr:    addr,
		MACAddr: "TODO",
	}
}
