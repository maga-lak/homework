package request

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"paswword"`
}

//todo валидация посмотреть validator
