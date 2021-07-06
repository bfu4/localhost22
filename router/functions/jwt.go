package functions

import (
	"cdn/util"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func ParseJWT(request *http.Request) (int, error) {
	cookie, err := request.Cookie("token")

	if err != nil {
		return -1, err
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return util.GetJWTSecret(), nil
	})

	if err != nil || !token.Valid {
		return -1, err
	}

	fmt.Printf("%+v\n", token)

	return 1, nil
}
