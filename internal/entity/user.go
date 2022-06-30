package entity

import "github.com/golang-jwt/jwt/v4"

type User struct {
	ID       string
	Username string
	Password string
	Name     string
}

// todo где расположить
type UserClaims struct {
	*jwt.RegisteredClaims
	User User `json:"user"`
}
