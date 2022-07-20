package main

import (
	"fmt"
	"log"
	"net/http"

	"shipple/bshipple/pkg/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	// upgrade the connection to a WebSocket connection
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	player := &websocket.Player{
		Conn: ws,
		Pool: pool,
	}

	// TODO when pool is full
	pool.Register <- player
	player.Read()
}

func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsEndpoint(pool, w, r)
	})
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}