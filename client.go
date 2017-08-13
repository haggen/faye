package faye

import (
	"errors"
	"time"
)

// Enumeration of Client states.
const (
	Unconnected = iota
	Connecting
	Connected
	Disconnected
)

// Client is the end-user interface with the protocol.
type Client struct {
	advice        *Advice
	state         int
	id            string
	transport     Transport
	extensions    []Extension
	subscriptions map[string]chan *Message
	incoming      chan *Message
	outgoing      chan *Message
}

// New initializes and returns a new Client.
func New(transport Transport) *Client {
	advice := &Advice{
		Interval:  time.Second,
		Timeout:   time.Second * 10,
		Reconnect: HandshakeAdvice,
	}
	c := &Client{
		advice:    advice,
		state:     Unconnected,
		transport: transport,
	}
	return c
}

// receive ...
func (c *Client) receive() (*Message, error) {
	m, err := c.transport.Receive()
	if err != nil {
		return nil, err
	}
	for _, ext := range c.extensions {
		ext.In(m)
	}
	return m, nil
}

// send ...
func (c *Client) send(m *Message) error {
	for _, ext := range c.extensions {
		ext.Out(m)
	}
	err := c.transport.Send(m)
	return err
}

// Handshake ...
func (c *Client) Handshake() error {
	if c.state != Connecting {
		return errors.New("faye: can ")
	}
	m := &Message{
		Version:                  Version,
		Channel:                  Handshake,
		SupportedConnectionTypes: []string{c.transport.ConnectionType()},
	}
	return c.send(m)
}

// Connect ...
func (c *Client) Connect() error {
	return nil
}

// Disconnect ...
func (c *Client) Disconnect() error {
	if c.state != Connected {
		return nil
	}
	m := &Message{
		ClientID: c.id,
		Channel:  Disconnect,
	}
	err := c.send(m)
	if err == nil {
		return err
	}
	return c.transport.Disconnect()
}
