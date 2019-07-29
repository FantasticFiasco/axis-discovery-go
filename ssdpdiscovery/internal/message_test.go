package ssdpdiscovery

import (
	"net"
	"testing"
)

func TestParseNotifyMessage(t *testing.T) {
	want := struct {
		addr, mac string
	}{
		"192.168.1.102", "ACCC8E270AD8",
	}
	got := message{ addr, []byte(notifyMessage)}.parse()
	if got.Addr != want.addr {
		t.Errorf("got.Addr == %q, want %q", got.Addr, want.addr)
	}
	if got.MACAddr != want.mac {
		t.Errorf("got.MACAddr == %q, want %q", got.MACAddr, want.mac)
	}
}

var addr = net.UDPAddr{
	IP:   []byte{192, 168, 1, 102},
	Port: 80,
	Zone: "",
}

var notifyMessage = "NOTIFY * HTTP/1.1\r\n" +
	"HOST: 239.255.255.250:1900\r\n" +
	"CACHE-CONTROL: max-age=1800\r\n" +
	"LOCATION: http://192.168.1.102:45895/rootdesc1.xml\r\n" +
	"OPT: \"http://schemas.upnp.org/upnp/1/0/\"; ns=01\r\n" +
	"01-NLS: 2ae7b584-1dd2-11b2-988f-983991d749b2\r\n" +
	"NT: urn:axis-com:service:BasicService:1\r\n" +
	"NTS: ssdp:byebye\r\n" +
	"SERVER: Linux/2.6.35, UPnP/1.0, Portable SDK for UPnP devices/1.6.18\r\n" +
	"X-User-Agent: redsonic\r\n" +
	"USN: uuid:Upnp-BasicDevice-1_0-ACCC8E270AD8::urn:axis-com:service:BasicService:1\r\n"