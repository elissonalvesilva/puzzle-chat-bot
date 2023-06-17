package main

import (
	"context"
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/websocket"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/pkg"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
)

func main() {
	m := pkg.NewMongoClient(os.Getenv("DB_HOST"))
	ctxTodo := context.TODO()
	client, err := m.Client()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctxTodo)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctxTodo)
	err = client.Ping(ctxTodo, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	app := db.NewDatabase(client, "test")

	ws := websocket.NewWebSocket(*app)
	http.HandleFunc("/ws", ws.WebsocketHandler)
	fmt.Println("WebSocket Starting")
	log.Fatal(http.ListenAndServe(":4513", nil))
}
