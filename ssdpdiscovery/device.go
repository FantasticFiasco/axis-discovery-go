package ssdpdiscovery

import "net"

// Device represents a network attached hardware unit from Axis Communications.
type Device struct {
	// Addr specifies the TCP address.
	Addr *net.TCPAddr

	// MACAddr is in most situations identical to the serial number. The exceptions are the Axis products which bundle
	// multiple physical devices into a single casing with a shared network interface. Because of the shared network
	// interface they also share the same MAC address.
	MACAddr string

	// FriendlyName specifies the short description of the device.
	FriendlyName string

	// ModelName specifies the short model name.
	ModelName string

	// ModelDescription specifies the long model description.
	ModelDescription string

	// ModelNumber is the short model number.
	ModelNumber string

	// PresentationURL specifies URL to the web page of the device.
	PresentationURL string
}
