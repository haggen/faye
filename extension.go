package faye

// Extension is an interface to mutate both incoming and outgoing messages.
type Extension interface {
	In(*Message)
	Out(*Message)
}
