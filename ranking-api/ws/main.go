package main

import (
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/websocket"
	"log"
	"net/http"
	"sync"
)

func WebSocketServer(db *db.Database) {
	ws := websocket.NewWebSocket(db)
	http.HandleFunc("/ws", ws.WebsocketHandler)
	fmt.Println("WebSocket Starting")
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	panic("Falha ao carregar o arquivo .env")
	//}
	app, err := db.NewDB()
	if err != nil {
		panic("Falha ao inicializar o aplicativo")
	}

	err = app.AutoMigrateTables()
	if err != nil {
		panic("Falha ao inicializar o migration")
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		WebSocketServer(app)
	}()

	wg.Wait()
}
