package models

import (
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Id       int    `json:"Id"`
	jwt.StandardClaims
}
