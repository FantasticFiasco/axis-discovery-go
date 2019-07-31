package ssdpdiscovery

import (
	"github.com/pkg/errors"
	"net"
	"regexp"
)

// Regexp looking for a MAC address in a notify message, displayed in the following format:
// "USN: uuid:Upnp-BasicDevice-1_0-0123456789AB::urn:axis-com:service:BasicService:1\r\n"
var macAddressFromMessageRegexp = regexp.MustCompile(
	"(?im)"				+ // Compiler flags, "i" for case insensitive and "m" for multiline
	"^"						+ // At beginning of line, since compiler flag "m" is set
	"USN:"					+ // The header name
	"\\s*"					+
	"uuid:"					+ // UUID prefix
	".*"					+
	"([0-9a-f]{12})"		+ // The MAC address
	"::"					+
	".*"					+
	"$")					  // At end of line, since compiler flag "m" is set

type message struct {
	addr	*net.UDPAddr
	b		[]byte
}

func newMessage(addr *net.UDPAddr, b []byte) *message {
	return &message{
		addr,
		b,
	}
}

func (m *message) parseNotify() (d Device, err error) {
	match := macAddressFromMessageRegexp.FindSubmatch(m.b)
	if match == nil  {
		err = errors.Errorf("MAC address not found in notify message %q", m.b)
		return
	}
	d = Device{
		Addr:             m.addr,
		MACAddr:          string(match[1]),
	}
	return
}
