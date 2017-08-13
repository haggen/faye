package faye

import (
	"time"
)

// Enumeration of reconnect advices.
const (
	HandshakeAdvice = "handshake"
	Retry           = "retry"
	None            = "none"
)

// Advice about what the server requires from the client.
type Advice struct {
	// Interval in milliseconds between connection checks.
	Interval time.Duration `json:"interval"`

	// Timeout in milliseconds for requests.
	Timeout time.Duration `json:"timeout"`

	// Advice string. May be 'none', 'handshake', or 'retry'.
	Reconnect string `json:"reconnect"`
}
