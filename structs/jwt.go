package structs

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserId int
}
