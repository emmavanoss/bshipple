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
    Ready bool
    Locations []gamestate.Location
}

type PlayerReadyMessage struct {
  Player *Player
}

type PlayerFireMessage struct {
  Player *Player
  Coordinates gamestate.Location
}

type Message struct {
  Type int
  Body string
}

type Input struct {
  Message string
  PlayerId string
  Battleships [][]int
  Coordinate []int
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

        // decode the JSON
        var decodedBody Input
        err = json.Unmarshal(p, &decodedBody)

        switch decodedBody.Message {
        case "start":
          c.ID = decodedBody.PlayerId

          locations := make([]gamestate.Location, 2)
          locations[0] = gamestate.Location{X: decodedBody.Battleships[0][0], Y: decodedBody.Battleships[0][1]}
          locations[1] = gamestate.Location{X: decodedBody.Battleships[1][0], Y: decodedBody.Battleships[0][1]}

          // update player state
          c.Ready = true
          c.Locations = locations

          // send player to the PlayerStart channel
          message := PlayerReadyMessage{Player: c}
          c.Pool.PlayerReady <- message
          // pp.Print(message)
        case "fire":
          coordinates := gamestate.Location{X: decodedBody.Coordinate[0], Y: decodedBody.Coordinate[1]}

          message := PlayerFireMessage{Player: c, Coordinates: coordinates}
          c.Pool.PlayerFire <- message
        default:
          return
        }
	}
}
