package main

import (
	"fmt"
	"github.com/FantasticFiasco/axis-discovery-go/ssdpdiscovery"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Listen passively for devices...")
	err := ssdpdiscovery.ListenPassive(alive, byeBye)
	if err != nil {
		panic(errors.Cause(err))
	}
}

func alive(d ssdpdiscovery.Device) {
	fmt.Printf("Received alive notification from MAC %s on address %s\n", d.MACAddr, d.Addr)
}

func byeBye(d ssdpdiscovery.Device) {
	fmt.Printf("Received byebye notification from MAC %s on address %s\n", d.MACAddr, d.Addr)
}
