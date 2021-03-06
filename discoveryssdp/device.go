package discoveryssdp

// A Device represents a device on the network.
type Device struct {
	Address				string	// network address
	Port				int		// network port
	MACAddress			string	// In most situations this is identical to the serial number. The exceptions are the Axis products which bundle multiple physical devices into a single casing with a shared network interface. Because of the shared network interface they also share the same MAC address.
	FriendlyName		string	// short description of the device
	ModelName			string	// short model name
	ModelDescription	string	// long model description
	ModelNumber			string	// short model number
	PresentationURL		string	// URL to the web page of the device
}
