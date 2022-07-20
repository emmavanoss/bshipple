package websocket

import (
	"fmt"
	"log"
)

type Pool struct {
	// TODO GameStateManager struct of some sort
	// TODO should be some kind of GameState message, not Message
	GameStateChange chan Message
	Players         []*Player
	PlayerAction    chan Message
	Register        chan *Player
}

func NewPool() *Pool {
	return &Pool{
		GameStateChange: make(chan Message),
		Players:         make([]*Player, 0, 2), // current length 0, max capacity 2
		PlayerAction:    make(chan Message),
		Register:        make(chan *Player),
	}
}

func (pool *Pool) Start() {
	// listen for messages passed to channels
	for {
		select {
		case newPlayer := <-pool.Register:
			fmt.Println("Registering new player to the pool")
			// TODO if len(Players) == 2, pool is full - return error?

			// Add new player to the pool
			pool.Players = append(pool.Players, newPlayer)

			fmt.Println("Number of players", len(pool.Players))

			break
		case message := <-pool.PlayerAction:
			fmt.Println("Handing PlayerAction")

			// TODO pass player action to GameState
			log.Println(message)

			break
		case message := <-pool.GameStateChange:
			fmt.Println("Handling GameStateChange")

			// pass game state change to Players
			for _, player := range pool.Players {
				player.Conn.WriteJSON(Message{Type: 1, Body: message.Body})
			}

			break
		}
	}
}
