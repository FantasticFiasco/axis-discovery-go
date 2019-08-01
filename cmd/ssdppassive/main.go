package main

import (
	"fmt"
	"github.com/FantasticFiasco/axis-discovery-go/ssdpdiscovery"
)

func main() {
	fmt.Println("Listen passively for devices...")
	err := ssdpdiscovery.ListenPassive(onAlive, onByeBye)
	if err != nil {
		panic(err)
	}
}

func onAlive(d ssdpdiscovery.Device) {
	fmt.Printf("Received alive notification from MAC %s on address %s\n", d.MACAddr, d.Addr)
}

func onByeBye(d ssdpdiscovery.Device) {
	fmt.Printf("Received byebye notification from MAC %s on address %s\n", d.MACAddr, d.Addr)
}
