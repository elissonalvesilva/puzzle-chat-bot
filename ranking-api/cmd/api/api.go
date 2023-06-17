package api

import (
	"encoding/json"
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/api/protocols"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"net/http"
)

type API struct {
	database db.MongoDatabase
}

type UserPostParam struct {
	Name    string `json:"name"`
	Current int    `json:"current"`
	Phone   string `json:"phone"`
}

type UserUpdateCurrent struct {
	Current int    `json:"current"`
	Phone   string `json:"phone"`
}

func NewAPI(db db.MongoDatabase) *API {
	return &API{
		database: db,
	}
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	var user UserPostParam
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Err to decode", http.StatusBadRequest)
		return
	}

	if _, err := api.database.GetByPhone(user.Phone); err == nil {
		http.Error(w, "Usuário Já cadastrado", http.StatusBadRequest)
		return
	}

	userToCreate := protocols.UserPostParam{
		Phone:   user.Phone,
		Name:    user.Name,
		Current: user.Current,
	}

	if err = api.database.Create(userToCreate); err != nil {
		fmt.Println(err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func (api *API) Update(w http.ResponseWriter, r *http.Request) {
	var user UserUpdateCurrent
	json.NewDecoder(r.Body).Decode(&user)

	userRanking, err := api.database.GetByPhone(user.Phone)
	if err != nil {
		http.Error(w, "Usuário nao cadastrado", http.StatusNotFound)
		return
	}

	userToUpdate := protocols.UserPostParam{
		Phone:   userRanking.Phone,
		Current: user.Current,
		Name:    userRanking.Name,
	}
	if err := api.database.Update(userRanking.Id, userToUpdate); err != nil {
		fmt.Println(err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func (api *API) Clean(w http.ResponseWriter, r *http.Request) {
	if err := api.database.DeleteAll(); err != nil {
		fmt.Println(err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func (api *API) Ranking(w http.ResponseWriter, r *http.Request) {
	users, err := api.database.GetAll()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
