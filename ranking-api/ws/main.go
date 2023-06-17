package main

import (
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/websocket"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/pkg"
	"log"
	"net/http"
	"os"
)

func main() {
	client, err := pkg.NewMongoClient(os.Getenv("DB_HOST")).Client()
	if err != nil {
		panic("Falha ao conectar ao db")
	}
	app := db.NewDatabase(client, "test")

	ws := websocket.NewWebSocket(*app)
	http.HandleFunc("/ws", ws.WebsocketHandler)
	fmt.Println("WebSocket Starting")
	log.Fatal(http.ListenAndServe(":4513", nil))
}
