// Package ssdpdiscovery provides means to discover network attached devices from Axis Communications using SSDP.
package ssdpdiscovery

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
)

// The multicast address where SSDP notification are announced
var multicastAddr = net.UDPAddr{
	IP:   []byte{239, 255, 255, 250},
	Port: 1900,
}

// ListenPassive will passively listen for SSDP notifications on the network.
func ListenPassive(onAlive func(d Device), onByeBye func(d Device)) error {
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
		fmt.Println(">>>")
		fmt.Printf("%+v", string(b[:n]))
		fmt.Println("---")
		fmt.Printf("%+v", m)
		fmt.Println("<<<")

		if m[method] != "NOTIFY * HTTP/1.1" ||
			m[nt] != "urn:axis-com:service:BasicService:1" {
			continue
		}

		if m[nts] == "ssdp:alive" {
			fmt.Println("ALIVE")
			onAlive(toDevice(addr, m))
		} else if m[nts] == "ssdp:byebye" {
			fmt.Println("BYEBYE")
			onByeBye(toDevice(addr, m))
		} else {
			fmt.Println("UNSUPPORTED: " + m[nts] + "\n\n")
		}

	}
}

func toDevice(addr *net.UDPAddr, m message) Device {
	return Device{
		//Addr:    addr,
		MACAddr: "TODO",
	}
}
