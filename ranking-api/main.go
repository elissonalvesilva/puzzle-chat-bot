package main

import (
	"fmt"
	api2 "github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"

	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
)

func WebAPI(db *db.Database) {
	router := mux.NewRouter()
	api := api2.NewAPI(db)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}).Methods("GET")
	router.HandleFunc("/api/create", api.Create).Methods("POST")
	router.HandleFunc("/api/update", api.Update).Methods("PUT")
	router.HandleFunc("/api/delete", api.Clean).Methods("DELETE")
	fmt.Println("API STartign")
	log.Fatal(http.ListenAndServe(":8002", router))
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
		WebAPI(app)
	}()

	wg.Wait()
}
