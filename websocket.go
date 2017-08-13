package faye

import (
	"github.com/gorilla/websocket"
)

// WebSocket connects and communicates with the server via WebSocket.
type WebSocket struct {
	url  string
	conn *websocket.Conn
}

// NewWebSocket initializes and returns a new WebSocket transport.
func NewWebSocket(url string) *WebSocket {
	return &WebSocket{
		url: url,
	}
}

// ConnectionType returns "websocket".
// This will be used when negoticating with the server the supported connection types.
func (ws *WebSocket) ConnectionType() string {
	return "websocket"
}

// Connect opens a connection to the WebSocket server.
func (ws *WebSocket) Connect() error {
	var err error
	ws.conn, _, err = websocket.DefaultDialer.Dial(ws.url, nil)
	return err
}

// Disconnect closes the connection to the WebSocket server.
func (ws *WebSocket) Disconnect() error {
	return nil
}

// Send writes a message encoded in JSON to the upstream.
func (ws *WebSocket) Send(m *Message) error {
	return ws.conn.WriteJSON(m)
}

// Receive reads a message encoded in JSON from the downstream.
func (ws *WebSocket) Receive() (*Message, error) {
	m := []*Message{}
	err := ws.conn.ReadJSON(&m)
	if err != nil {
		return nil, err
	}
	return m[0], nil
}

// Ping write a ping message to the WebSocket server.
func (ws *WebSocket) Ping() error {
	return ws.conn.WriteMessage(websocket.PingMessage, []byte("[]"))
}
