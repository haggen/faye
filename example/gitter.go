package main

import (
	"log"

	"github.com/haggen/faye"
)

type Auth struct {
	token string
}

func (a *Auth) In(m *faye.Message) {}

func (a *Auth) Out(m *faye.Message) {
	if m.Channel == faye.Handshake {
	}
}

func main() {
	ws, err := faye.NewWebSocket("wss://ws.gitter.im/faye")
	f, err := faye.NewClient(ws)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Disconnect()
	f.AddExtension(&Auth{"3b44f81a7a49e8252f5f9ebfd19b3c3a81fc003c"})
	channel := f.Subscribe(`/api/v1/rooms/56a7e44de610378809be5844/events`)
	for {
		m := <-channel
		log.Printf("%+v", m)
	}
}

// func main() {
// 	header := http.Header{}
// 	header.Set("Content-type", "application/json")
// 	conn, _, err := websocket.DefaultDialer.Dial(`wss://transport.gitter.im/faye`, header)
// 	if err != nil {
// 		log.Fatalln("Dial:", err)
// 	}
// 	defer conn.Close()
// 	go func() {
// 		for {
// 			var m []*Message
// 			err := conn.ReadJSON(&m)
// 			if err != nil {
// 				log.Printf("ReadJSON: %+v", err)
// 			} else {
// 				log.Printf("ReadJSON: %+v", m[0])
// 			}
// 		}
// 	}()
// 	ext := struct {
// 		Token string `json:"token"`
// 	}{
// 		Token: "",
// 	}
// 	m := &Message{
// 		Ext:                      ext,
// 		Version:                  Version,
// 		Channel:                  Handshake,
// 		SupportedConnectionTypes: []string{"websocket"},
// 	}
// 	err = conn.WriteJSON(m)
// 	if err != nil {
// 		log.Println("WriteJSON:", err)
// 	}
// 	block := make(chan int, 1)
// 	<-block
// }
