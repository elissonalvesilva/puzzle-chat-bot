package main

import (
	"log"
	"net/http"
)

import (
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir todas as origens (INSECURE)
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Agora você pode ler e escrever mensagens usando a conexão WebSocket
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// Processar a mensagem recebida

		// Enviar uma mensagem de volta para o cliente
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", websocketHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
