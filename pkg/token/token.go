package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var key = []byte("Better_Dorm_@CCNU")

type Myclaims struct {
	UID int
	jwt.StandardClaims
}

func Newtoken(UID int) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, Myclaims{
		UID,
		jwt.StandardClaims{
			Issuer:    "Muxi_Studio",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		}}).SignedString(key)
}

func Parsetoken(tokenString string) (*jwt.Token, *Myclaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		return nil, nil, err
	}
	cl, _ := token.Claims.(*Myclaims)
	return token, cl, err
}
