package api

import (
	"encoding/json"
	"fmt"
	"github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"
	"net/http"
)

type API struct {
	database *db.Database
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

func NewAPI(db *db.Database) *API {
	return &API{
		database: db,
	}
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	var user UserPostParam
	json.NewDecoder(r.Body).Decode(&user)

	if ok := api.database.ExistsPhone(user.Phone); ok {
		http.Error(w, "Usuário Já cadastrado", http.StatusBadRequest)
		return
	}

	userToCreate := db.Ranking{
		Phone:   user.Phone,
		Name:    user.Name,
		Current: user.Current,
	}

	if err := api.database.Create(userToCreate); err != nil {
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

	userToUpdate := db.Ranking{
		ID:      userRanking.ID,
		Phone:   userRanking.Phone,
		Current: user.Current,
		Name:    userRanking.Name,
	}
	if err := api.database.UpdateCurrent(userToUpdate); err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *API) Clean(w http.ResponseWriter, r *http.Request) {
	if err := api.database.DeleteAll(); err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
