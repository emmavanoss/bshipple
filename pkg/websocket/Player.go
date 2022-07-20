package websocket

import (
	"log"
    "encoding/json"

	"github.com/gorilla/websocket"
    "shipple/bshipple/pkg/gamestate"
    // "github.com/k0kubun/pp"
)

type Player struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type PlayerReadyMessage struct {
  Player *Player
  Locations []gamestate.Location
}

type Message struct {
  Type int
  Body string
}

type StartInput struct {
  Message string
  PlayerId string
  Battleships [][]int
}



// Read from the websocket & broadcast to PlayerAction channel
func (c *Player) Read() {
	defer func() {
		// close conn if there's an error, maybe?
		c.Conn.Close()
	}()

	for {
		// receive message from the websocket connection
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

        // TODO more than just start input is possible
        var decodedBody StartInput
        err = json.Unmarshal(p, &decodedBody)

        if decodedBody.Message == "start" {
          c.ID = decodedBody.PlayerId

          locations := make([]gamestate.Location, 2)
          locations[0] = gamestate.Location{X: decodedBody.Battleships[0][0], Y: decodedBody.Battleships[0][1]}
          locations[1] = gamestate.Location{X: decodedBody.Battleships[1][0], Y: decodedBody.Battleships[0][1]}

          // send player and locations to the PlayerStart channel
          message := PlayerReadyMessage{Player: c, Locations: locations}
          c.Pool.PlayerReady <- message
          // pp.Print(message)
        }

        // TODO  fire
	}
}
