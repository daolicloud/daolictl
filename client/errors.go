package client

import (
	"errors"
)

// ErrConnectionFailed is a error raised when the connection between the client and the server failed.
var ErrConnectionFailed = errors.New("Cannot connect to the Daolinet server. Is the daolinet server running on this host?")
