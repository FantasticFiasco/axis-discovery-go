package ssdpdiscovery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMethod(t *testing.T) {
	for _, test := range []struct {
		in   []byte
		want string
	}{
		{notifyAliveMessage, "NOTIFY * HTTP/1.1"},
		{notifyByeByeMessage, "NOTIFY * HTTP/1.1"},
		{mSearchMessage, "HTTP/1.1 200 OK"},
	} {
		got := parseMessage(test.in).method
		want := test.want
		assert.Equal(t, want, got)
	}
}

func TestParseLocation(t *testing.T) {
	for _, test := range []struct {
		in []byte
	}{
		{notifyAliveMessage},
		{notifyByeByeMessage},
		{mSearchMessage},
	} {
		got := parseMessage(test.in).location
		want := "http://192.168.1.102:45895/rootdesc1.xml"
		assert.Equal(t, want, got)
	}
}

func TestParseUSN(t *testing.T) {
	for _, test := range []struct {
		in []byte
	}{
		{notifyAliveMessage},
		{notifyByeByeMessage},
		{mSearchMessage},
	} {
		got := parseMessage(test.in).usn
		want := "uuid:Upnp-BasicDevice-1_0-ACCC8E270AD8::urn:axis-com:service:BasicService:1"
		assert.Equal(t, want, got)
	}
}

func TestParseNT(t *testing.T) {
	for _, test := range []struct {
		in []byte
	}{
		{notifyAliveMessage},
		{notifyByeByeMessage},
	} {
		got := parseMessage(test.in).nt
		want := "urn:axis-com:service:BasicService:1"
		assert.Equal(t, want, got)
	}
}

func TestParseNTS(t *testing.T) {
	for _, test := range []struct {
		in   []byte
		want string
	}{
		{notifyAliveMessage, "ssdp:alive"},
		{notifyByeByeMessage, "ssdp:byebye"},
	} {
		got := parseMessage(test.in).nts
		want := test.want
		assert.Equal(t, want, got)
	}
}

var notifyAliveMessage = []byte(
	"NOTIFY * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"CACHE-CONTROL: max-age=1800\r\n" +
		"LOCATION: http://192.168.1.102:45895/rootdesc1.xml\r\n" +
		"OPT: \"http://schemas.upnp.org/upnp/1/0/\"; ns=01\r\n" +
		"01-NLS: 2ae7b584-1dd2-11b2-988f-983991d749b2\r\n" +
		"NT: urn:axis-com:service:BasicService:1\r\n" +
		"NTS: ssdp:alive\r\n" +
		"SERVER: Linux/2.6.35, UPnP/1.0, Portable SDK for UPnP devices/1.6.18\r\n" +
		"X-User-Agent: redsonic\r\n" +
		"USN: uuid:Upnp-BasicDevice-1_0-ACCC8E270AD8::urn:axis-com:service:BasicService:1\r\n")

var notifyByeByeMessage = []byte(
	"NOTIFY * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"CACHE-CONTROL: max-age=1800\r\n" +
		"LOCATION: http://192.168.1.102:45895/rootdesc1.xml\r\n" +
		"OPT: \"http://schemas.upnp.org/upnp/1/0/\"; ns=01\r\n" +
		"01-NLS: 2ae7b584-1dd2-11b2-988f-983991d749b2\r\n" +
		"NT: urn:axis-com:service:BasicService:1\r\n" +
		"NTS: ssdp:byebye\r\n" +
		"SERVER: Linux/2.6.35, UPnP/1.0, Portable SDK for UPnP devices/1.6.18\r\n" +
		"X-User-Agent: redsonic\r\n" +
		"USN: uuid:Upnp-BasicDevice-1_0-ACCC8E270AD8::urn:axis-com:service:BasicService:1\r\n")

var mSearchMessage = []byte(
	"HTTP/1.1 200 OK\r\n" +
		"CACHE-CONTROL: max-age=1800\r\n" +
		"DATE: Sun, 02 Oct 2016 21:11:25 GMT\r\n" +
		"EXT:\r\n" +
		"LOCATION: http://192.168.1.102:45895/rootdesc1.xml\r\n" +
		"OPT: \"http://schemas.upnp.org/upnp/1/0/\"; ns=01\r\n" +
		"01-NLS: 8fb2638a-1dd2-11b2-a915-c89968cce2ca\r\n" +
		"SERVER: Linux/2.6.35, UPnP/1.0, Portable SDK for UPnP devices/1.6.18\r\n" +
		"X-User-Agent: redsonic\r\n" +
		"ST: urn:axis-com:service:BasicService:1\r\n" +
		"USN: uuid:Upnp-BasicDevice-1_0-ACCC8E270AD8::urn:axis-com:service:BasicService:1\r\n")
