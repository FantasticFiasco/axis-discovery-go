// Package ssdpdiscovery provides means to discover devices from Axis Communications on the network using SSDP.
package ssdpdiscovery

import (
	"fmt"
	"net"
)

// ListenPassive todo
func ListenPassive() (err error) {
	multicastAddress := net.UDPAddr{
		IP:   []byte{239, 255, 255, 250},
		Port: 1900,
		Zone: "",
	}
	conn, err := net.ListenMulticastUDP("udp", nil, &multicastAddress)
	if err != nil {
		return
	}
	defer conn.Close()
	b := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(b)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("Received %s from %s\n", string(b[0:n]), addr)
		}
	}
}
