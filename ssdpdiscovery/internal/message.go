package ssdpdiscovery

import (
	"errors"
	"fmt"
	"github.com/FantasticFiasco/axis-discovery-go/ssdpdiscovery"
	"net"
	"regexp"
)

// Regexp looking for a MAC address in a notify message, displayed in the following format:
// "USN: uuid:Upnp-BasicDevice-1_0-0123456789AB::urn:axis-com:service:BasicService:1\r\n"
var macAddressFromMessageRegexp = regexp.MustCompile(
	"(?im)"				+ // Compiler flags, "i" for case insensitive and "m" for multiline
	"^"						+ // At beginning of line, since compiler flag "m" is set
	"USN:"					+ // The header name
	"\\s*"					+ // Zero or more whitespaces
	"uuid:"					+ // UUID prefix
	".*"					+ // Zero or more any characters
	"([0-9a-f]{12})"		+ // The MAC address
	"::"					+ // Separator
	".*"					+ // Zero or more any characters
	"$")					  // At end of line, since compiler flag "m" is set

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
	match := macAddressFromMessageRegexp.FindSubmatch(m.b)
	if match == nil  {
		err = errors.New(fmt.Sprintf("MAC address not found in notify message %q", m.b))
		return
	}
	d = ssdpdiscovery.Device{
		Addr:             m.addr,
		MACAddr:          string(match[1]),
	}
	return
}
