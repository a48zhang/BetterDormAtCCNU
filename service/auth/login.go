package auth

import (
	"context"
	"errors"
	"main/model"
	"main/pkg/token"
)

func Login(ctx context.Context, LoginInfo *model.User) (string, error) {
	loginName := LoginInfo.Name
	loginPwd := LoginInfo.Passwd
	dbinfo := model.User{Name: loginName}
	err := dbinfo.FindByName(ctx)
	if err != nil {
		return "", err
	}
	if dbinfo.Passwd != loginPwd {
		return "", errors.New("wrong username or password")
	} else {
		return token.Newtoken(dbinfo.CCNUid)
	}
}
