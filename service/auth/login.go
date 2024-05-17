package auth

import (
	"context"
	"errors"
	"main/dao"
	"main/pkg/token"
)

func Login(ctx context.Context, LoginInfo *dao.User) (string, error) {
	loginName := LoginInfo.Name
	loginPwd := LoginInfo.Passwd
	dbinfo := dao.User{Name: loginName}
	err := dbinfo.FindByName(ctx)
	if err != nil {
		return "", err
	}
	if dbinfo.Passwd != loginPwd {
		return "", errors.New("wrong username or password")
	} else {
		return token.NewTokenForUser(dbinfo)
	}
}
