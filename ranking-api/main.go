package main

import (
	api2 "github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"

	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/websocket"
)

func WebSocketServer(db *db.Database) {
	ws := websocket.NewWebSocket(db)

	http.HandleFunc("/ws", ws.WebsocketHandler)
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

	router := mux.NewRouter()

	api := api2.NewAPI(app)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		WebSocketServer(app)
	}()
	wg.Wait()

	router.HandleFunc("/api/create", api.Create).Methods("POST")
	router.HandleFunc("/api/update", api.Create).Methods("PUT")
	router.HandleFunc("/api/delete", api.Create).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8002", router))
}
