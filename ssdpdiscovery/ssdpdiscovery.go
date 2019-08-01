// Package ssdpdiscovery provides means to discover devices from Axis Communications on the network using SSDP.
package ssdpdiscovery

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
)

var multicastAddress = net.UDPAddr{
	IP:   []byte{239, 255, 255, 250},
	Port: 1900,
}

func ListenPassive(onAlive func(d Device), onByeBye func(d Device)) error {
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
		m := parseMessage(b[:n])
		if m.method != "NOTIFY * HTTP/1.1" ||
			m.nt != "urn:axis-com:service:BasicService:1" {
			continue
		}

		if m.nts == "ssdp:alive" {
			fmt.Println("ALIVE")
			onAlive(toDevice(addr, m))
		} else if m.nts == "ssdp:byebye" {
			fmt.Println("BYEBYE")
			onByeBye(toDevice(addr, m))
		} else {
			fmt.Println("UNSUPPORTED: " + m.nts + "\n\n")
		}

	}
}

func toDevice(addr *net.UDPAddr, m message) Device {
	return Device{
		Addr:    addr,
		MACAddr: "TODO",
	}
}
