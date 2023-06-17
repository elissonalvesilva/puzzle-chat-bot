package protocols

import "github.com/elissonalvesilva/puzzle-chat-bot/ranking-api/cmd/db"

type UserPostParam struct {
	Name    string `json:"name"`
	Current int    `json:"current"`
	Phone   string `json:"phone"`
}

type UserGetResponse struct {
	Id string `json:"id"`
	UserPostParam
}

type UsersRanking struct {
	Users []db.UserModel
}

type UserUpdateCurrent struct {
	Current int    `json:"current"`
	Phone   string `json:"phone"`
}
