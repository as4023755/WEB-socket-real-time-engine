package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/as4023755/WEB-socket-real-time-engine/internal/websocket"
)

func main() {
	fmt.Println("MAIN STARTED")
	hub := websocket.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
