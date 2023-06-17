package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/gorilla/websocket"
	"net/http"
)

type WS struct {
	clients map[*websocket.Conn]bool
	db      *db.Database
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir todas as origens (INSECURE)
	},
}

func NewWebSocket(db *db.Database) *WS {
	return &WS{
		clients: make(map[*websocket.Conn]bool),
		db:      db,
	}
}

func (ws *WS) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	fmt.Println(connection)
	ws.clients[connection] = true // Save the connection using it as a key

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Exit the loop if the client tries to close the connection or the connection is interrupted
		}

		if string(message) == "refresh" {
			ranking, err := ws.db.GetAll()
			if err != nil {
				fmt.Println(err)
			}

			response, _ := json.Marshal(ranking)
			//response := []byte("Mensagem de refresh recebida!")
			err = connection.WriteMessage(websocket.TextMessage, response)
			if err != nil {
				fmt.Println("Falha ao enviar mensagem de resposta:", err)
			}
		} else {
			go ws.handleMessage(message)
		}
	}

	delete(ws.clients, connection) // Removing the connection

	connection.Close()
}

func (ws *WS) handleMessage(message []byte) {
	fmt.Println(string(message))
}
