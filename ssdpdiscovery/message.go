package ssdpdiscovery

import (
	"bufio"
	"strings"
)

type message map[messageKey]string

type messageKey string

const (
	// Empty key is indicating "type", i.e. the only parameter not formatted as a key/value pair
	method   messageKey = ""
	location messageKey = "LOCATION"
	nt       messageKey = "NT"
	nts      messageKey = "NTS"
	usn      messageKey = "USN"
)

func parseMessage(b []byte) message {
	m := make(map[messageKey]string)
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		var key, value string
		if len(parts) == 1 {
			key = string(method)
			value = parts[0]
		} else {
			key = parts[0]
			value = strings.Join(parts[1:], ":")
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		m[messageKey(key)] = value
	}
	return m
}
