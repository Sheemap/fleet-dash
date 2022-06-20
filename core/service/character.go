package service

import "github.com/golang-jwt/jwt/v4"

type Character struct {
	ID string
	Name string
	AccessToken *jwt.Token
}