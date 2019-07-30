package ssdpdiscovery

import (
	"fmt"
	"github.com/FantasticFiasco/axis-discovery-go/ssdpdiscovery"
	"net"
	"regexp"
)

var notifyRegexp = regexp.MustCompile(`(?im)^USN: uuid:.*([0-9a-f]{12})::.*$`)

type message struct {
	addr	net.UDPAddr
	b		[]byte
}

func newMessage(addr net.UDPAddr, b []byte) *message {
	return &message{
		addr,
		b,
	}
}

func (m *message) parseNotify() (d ssdpdiscovery.Device, err error) {
	temp:= notifyRegexp.FindSubmatch(m.b)
	for _, v := range temp {
		fmt.Println(string(v))
	}

	d = ssdpdiscovery.Device{
		Addr:             m.addr,
		MACAddr:          string(temp[1]),
	}

	return
}
