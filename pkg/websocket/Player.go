package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// Read from the websocket & broadcast to PlayerAction channel
func (c *Player) Read() {
	defer func() {
		// close conn if there's an error, maybe?
		c.Conn.Close()
	}()

	for {
		// receive message from the websocket connection
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// send message to the PlayerAction channel
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.PlayerAction <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
