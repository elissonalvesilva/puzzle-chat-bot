package main

import (
	"fmt"
	api2 "github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/api"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func WebAPI(db *db.Database) {

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

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}).Methods("GET")
	router.HandleFunc("/api/create", api.Create).Methods("POST")
	router.HandleFunc("/api/update", api.Update).Methods("PUT")
	router.HandleFunc("/api/delete", api.Clean).Methods("DELETE")
	fmt.Println("API STartign")
	log.Fatal(http.ListenAndServe(":4546", router))
}
