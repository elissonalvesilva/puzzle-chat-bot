package protocols

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
	Users []UserGetResponse
}

type UserUpdateCurrent struct {
	Current int    `json:"current"`
	Phone   string `json:"phone"`
}
