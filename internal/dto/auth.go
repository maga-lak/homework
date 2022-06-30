package dto

type AuthView struct {
	Token    string `json:"token"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}
