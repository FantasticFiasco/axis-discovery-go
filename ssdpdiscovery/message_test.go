package ssdpdiscovery

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestParseMessageGivenMethod(t *testing.T) {
	for _, test := range []struct {
		in   []byte
		want method
	}{
		{notifyAliveMessage, notify},
		{notifyByeByeMessage, notify},
		{mSearchMessage, mSearch},
		{[]byte("ABC"), nop},
		{[]byte(""), nop},
	} {
		got := parseMessage(test.in).method
		want := test.want
		assert.Equal(t, want, got)
	}
}

func TestParseMessageGivenLocation(t *testing.T) {
	for _, test := range []struct {
		in []byte
	}{
		{notifyAliveMessage},
		{notifyByeByeMessage},
		{mSearchMessage},
	} {
		got := parseMessage(test.in).headers[location]
		want := "http://192.168.1.102:45895/rootdesc1.xml"
		assert.Equal(t, want, got)
	}
}

func TestParseMessageGivenNT(t *testing.T) {
	for _, test := range []struct {
		in   []byte
		want string
	}{
		{notifyAliveMessage, "urn:axis-com:service:BasicService:1"},
		{notifyByeByeMessage, "urn:axis-com:service:BasicService:1"},
		{mSearchMessage, ""},
	} {
		got := parseMessage(test.in).headers[nt]
		want := test.want
		assert.Equal(t, want, got)
	}
}

func TestParseMessageGivenNTS(t *testing.T) {
	for _, test := range []struct {
		in   []byte
		want string
	}{
		{notifyAliveMessage, "ssdp:alive"},
		{notifyByeByeMessage, "ssdp:byebye"},
		{mSearchMessage, ""},
	} {
		got := parseMessage(test.in).headers[nts]
		want := test.want
		assert.Equal(t, want, got)
	}
}

func TestParseMessageGivenUSN(t *testing.T) {
	for _, test := range []struct {
		in []byte
	}{
		{notifyAliveMessage},
		{notifyByeByeMessage},
		{mSearchMessage},
	} {
		got := parseMessage(test.in).headers[usn]
		want := "uuid:Upnp-BasicDevice-1_0-ACCC8E270AD8::urn:axis-com:service:BasicService:1"
		assert.Equal(t, want, got)
	}
}
