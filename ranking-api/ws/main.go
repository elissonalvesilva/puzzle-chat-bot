package main

import (
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/websocket"
	"log"
	"net/http"
)

func WebSocketServer(db *db.Database) {

}

func main() {
	app, err := db.NewDB()
	if err != nil {
		panic("Falha ao inicializar o aplicativo")
	}

	err = app.AutoMigrateTables()
	if err != nil {
		panic("Falha ao inicializar o migration")
	}

	ws := websocket.NewWebSocket(app)
	http.HandleFunc("/ws", ws.WebsocketHandler)
	fmt.Println("WebSocket Starting")
	log.Fatal(http.ListenAndServe(":4513", nil))
}
