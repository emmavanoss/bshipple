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
	PlayerReady    chan PlayerReadyMessage
	PlayerFire    chan PlayerFireMessage
	Register        chan *Player
}

func NewPool() *Pool {
	return &Pool{
		GameStateChange: make(chan Message),
		Players:         make([]*Player, 0, 2), // current length 0, max capacity 2
		PlayerReady:     make(chan PlayerReadyMessage),
		PlayerFire:      make(chan PlayerFireMessage),
		Register:        make(chan *Player),
	}
}

func (pool *Pool) Start() {
	// listen for messages passed to channels
	for {
		select {
		case newPlayer := <-pool.Register:
            if len(pool.Players) >= 2 {
              fmt.Println("Pool is full no more players in the pool")
              break
            }

			fmt.Println("Registering new player to the pool")

			// Add new player to the pool
			pool.Players = append(pool.Players, newPlayer)
			fmt.Println("Number of players", len(pool.Players))

            go newPlayer.Read()
			break
		case message := <-pool.PlayerReady:
			fmt.Println("Handling PlayerReady")

			// TODO pass to GameState
			log.Println(message)

			break
		case message := <-pool.PlayerFire:
			fmt.Println("Handling PlayerFire")

			// TODO pass to GameState
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
