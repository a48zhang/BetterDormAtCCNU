package token

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/dao"
	"time"

	"github.com/golang-jwt/jwt"
)

var key = []byte("Better_Dorm_@CCNU")

type Myclaims struct {
	UID      string
	Username string
	UserMID  primitive.ObjectID
	jwt.StandardClaims
}

func NewTokenForUser(u dao.User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, Myclaims{
		u.CCNUid,
		u.Name,
		u.MID,
		jwt.StandardClaims{
			Issuer:    "Muxi_Studio",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		}}).SignedString(key)
}

// Newtoken deprecated
func Newtoken(UID string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, Myclaims{
		UID,
		"",
		primitive.NilObjectID,
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
