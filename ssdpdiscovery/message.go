package ssdpdiscovery

import (
	"bufio"
	"strings"
)

type message struct {
	method   string
	location string
	usn      string
	nt       string
	nts      string
}

func parseMessage(b []byte) (m message) {
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	// First line is always the method, and is not formatted as the other lines
	if !scanner.Scan() {
		return
	}
	m.method = scanner.Text()
	// Enumerate key/value pairs
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) < 2 {
			continue
		}
		key := parts[0]
		value := strings.TrimSpace(strings.Join(parts[1:], ":"))
		switch key {
		case "LOCATION":
			m.location = value
		case "USN":
			m.usn = value
		case "NT":
			m.nt = value
		case "NTS":
			m.nts = value
		}
	}
	return
}
