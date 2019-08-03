package ssdpdiscovery

import "log"

// Logger is an internal debug logger. Normally there is no need to expose the internal log messages, but if
// there is a need for it the messages can be exposed by setting this variable.
var Logger *log.Logger

func logf(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf(format, v...)
	}
}
