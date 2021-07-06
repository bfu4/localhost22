package functions

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtWrapper struct {
	SecretKey       []byte
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	UserId int
	jwt.StandardClaims
}

type SignedTokenResult struct {
	token   string
	expires int64
}

func (j *JwtWrapper) GetExpiry() time.Time {
	return time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours))
}

// GenerateToken generates a jwt token
func (j *JwtWrapper) GenerateToken(userId int, expiry time.Time) (signedToken string, err error) {
	claims := &JwtClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString(j.SecretKey)

	if err != nil {
		return
	}

	return
}

//ValidateToken validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SecretKey, nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)

	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}

	return

}
