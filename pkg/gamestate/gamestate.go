package gamestate

import (
	"fmt"
	"golang.org/x/exp/slices"
    "shipple/bshipple/pkg/websocket"
)

type Location struct {
	X, Y int
}

type LocationState struct {
	Ship bool
	Hit bool
}

type Board struct {
	Locations map[Location]LocationState
}

type GameState struct {
	Player1 websocket.Player
	Board1 Board
	Player2 string
	Board2 websocket.Player
}

// run when both players are ready
func StartGame(players *[]websocket.Player) {
    player1 := players[0]
    player2 := players[1]

	board1 := createBoard(player1.Locations)
	board2 := createBoard(player2.Locations)

	state := &GameState{Player1: player1, Board1: *board1, Player2: player2, Board2: *board2}

	return state
}

func Fire(state *GameState, location *Location, player *Player) *GameState {
	if state.Player1 == player {
		state.Board2.Locations[*location] = LocationState{
			Ship: state.Board2.Locations[*location].Ship,
			Hit: true,
		}
	} else {
		state.Board1.Locations[*location] = LocationState{
			Ship: state.Board1.Locations[*location].Ship,
			Hit: true,
		}
	}

	return state
}

func IsWinner(state *GameState, player string) bool {
	if state.Player1 == player {
		for _, location := range state.Board2.Locations {
			if location.Ship && !location.Hit {
				return false
			}
		}
	} else {
		for _, location := range state.Board1.Locations {
			if location.Ship && !location.Hit {
				return false
			}
		}
	}

	return true
}

func createBoard(battleships *[]Location) *Board {
	var board Board
	board.Locations = make(map[Location]LocationState)

	for x := 1; x <= 4; x++ {
		for y := 1; y <= 4; y++ {
			if slices.Contains(*battleships, Location{x, y}) {
				board.Locations[Location{x,y}] = LocationState{Ship: true, Hit: false}
			} else {
				board.Locations[Location{x,y}] = LocationState{Ship: false, Hit: false}
			}
		}
	}

	return &board
}

func DoTheThing() {
	// board := make(map[gamestate.Location]gamestate.LocationState)
	// board[gamestate.Location{1, 1}] = gamestate.LocationState{Ship: true, Hit: false}

	locations1 := make([]Location, 2)
	locations1[0] = Location{1, 1}
	locations1[1] = Location{1, 2}

	locations2 := make([]Location, 2)
	locations2[0] = Location{3, 4}
	locations2[1] = Location{4, 4}

	// player1 = *"Player1"

	state := StartGame("Player1", &locations1, "Player2", &locations2)
	// fmt.Println(state)

	// fmt.Println(board)

	Fire(state, &Location{1, 1}, "Player2")
	Fire(state, &Location{1, 2}, "Player2")
	// newBoard := Fire(state, &Location{1, 1}, "Player1")
	// fmt.Println(newBoard)

	// var isWinner bool

	fmt.Println(IsWinner(state, "Player2"))
	fmt.Println(IsWinner(state, "Player1"))
}
