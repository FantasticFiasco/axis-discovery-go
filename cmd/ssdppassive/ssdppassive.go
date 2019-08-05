package main

import (
	"fmt"
	"github.com/FantasticFiasco/axis-discovery-go/ssdpdiscovery"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {
	fmt.Println("Listen passively for devices...")
	ssdpdiscovery.Logger = log.New(os.Stderr, "[Discovery(SSDP)] ", log.LstdFlags)
	err := ssdpdiscovery.Lurk(onAlive, onByeBye)
	if err != nil {
		panic(errors.Cause(err))
	}
}

func onAlive(d *ssdpdiscovery.Device) {
	fmt.Printf("Received alive notification from MAC %s on address %s\n", d.MACAddr, d.Addr)
}

func onByeBye(d *ssdpdiscovery.Device) {
	fmt.Printf("Received byebye notification from MAC %s on address %s\n", d.MACAddr, d.Addr)
}
