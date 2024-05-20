package auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"main/dao"
	"main/pkg/token"
)

func Login(ctx context.Context, name string, pwd string) (string, error) {
	loginName := name
	dbinfo := dao.User{Name: loginName}
	err := dbinfo.FindByName(ctx)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return "", errors.New("wrong username or password (error code:101)")
	}
	if err != nil {
		return "", err
	}
	if dbinfo.Passwd.Is(pwd) {
		return token.NewTokenForUser(dbinfo)
	} else {
		return "", errors.New("wrong username or password (error code:102)")
	}
}
