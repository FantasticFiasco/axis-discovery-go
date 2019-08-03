package ssdpdiscovery

import (
	"bufio"
	"strings"
)

type message struct {
	method  method
	headers map[headerName]string
}

type method string

const nop = ""

const (
	notify  method = "NOTIFY * HTTP/1.1"
	mSearch method = "HTTP/1.1 200 OK"
)

type headerName string

const (
	location headerName = "LOCATION"
	nt       headerName = "NT"
	nts      headerName = "NTS"
	usn      headerName = "USN"
)

func parseMessage(b []byte) *message {
	m := message{
		headers: make(map[headerName]string),
	}
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	// Parse method
	scanner.Scan()
	m.method = parseMessageMethod(scanner.Text())
	// Parse headers
	for scanner.Scan() {
		name, value := parseMessageHeader(scanner.Text())
		if name != "" {
			m.headers[name] = value
		}
	}
	return &m
}

func parseMessageMethod(s string) method {
	switch s {
	case string(notify):
		return notify
	case string(mSearch):
		return mSearch
	default:
		return nop
	}
}

func parseMessageHeader(s string) (name headerName, value string) {
	parts := strings.Split(s, ":")
	// Parse name
	switch parts[0] {
	case string(location):
		name = location
	case string(nt):
		name = nt
	case string(nts):
		name = nts
	case string(usn):
		name = usn
	}
	// Parse value
	value = strings.TrimSpace(strings.Join(parts[1:], ":"))
	return
}
