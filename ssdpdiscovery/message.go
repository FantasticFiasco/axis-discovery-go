package ssdpdiscovery

import (
	"bufio"
	"reflect"
	"strings"
)

type message struct {
	reflect.Method string
	map[headerName]string
}

type headerName string

const (
	// The first line in a message is the method, and it is not a header. Lets indicate this method as the header with
	// an empty string as its name.
	method   headerName = ""
	location headerName = "LOCATION"
	nt       headerName = "NT"
	nts      headerName = "NTS"
	usn      headerName = "USN"
)

func parseMessage(b []byte) message {
	m := make(map[headerName]string)
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
		m[headerName(key)] = value
	}
	return m
}

func (m *message) getHeaderValue(headerName headerName)
