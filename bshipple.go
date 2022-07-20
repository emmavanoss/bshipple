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

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
    ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

    go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}



	// err = ws.WriteMessage(1, []byte("Hi Client!"))
	// if err != nil {
		// log.Println(err)
	// }
