package faye

// Transport is the interface for communication with the server. e.g. WebSocket or Polling.
type Transport interface {
	Connect(string) error
	Disconnect() error
	Send(*Message) error
	Receive() (*Message, error)
	ConnectionType() string
	Ping()
}
