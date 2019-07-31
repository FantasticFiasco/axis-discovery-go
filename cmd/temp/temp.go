package main

import (
	"fmt"
	"github.com/FantasticFiasco/axis-discovery-go/ssdpdiscovery"
)

func main() {
	err := ssdpdiscovery.ListenPassive(log)
	if err != nil {
		panic(err)
	}
}

func log(d ssdpdiscovery.Device) {
	fmt.Printf("Received message from MAC %s on address %s\n", d.MACAddr, d.Addr.String())
}
