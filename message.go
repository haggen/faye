package faye

// Message encode/decode both requests and responses.
type Message struct {
	ID                       int         `json:"id"`
	Timestamp                int64       `json:"timestamp"`
	Ext                      interface{} `json:"ext"`
	Data                     interface{} `json:"data"`
	Advice                   *Advice     `json:"advice"`
	Version                  string      `json:"version"`
	ClientID                 string      `json:"clientId"`
	Channel                  string      `json:"channel"`
	Error                    string      `json:"error"`
	Successful               bool        `json:"successful"`
	Subscription             string      `json:"subscription"`
	SupportedConnectionTypes []string    `json:"supportedConnectionTypes"`
}
