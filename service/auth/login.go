package auth

import (
	"context"
	"errors"
	"main/model"
	"main/pkg/token"
)

func Login(LoginInfo *model.User) (string, error) {
	loginID := LoginInfo.Uid
	loginPwd := LoginInfo.Passwd
	dbinfo := model.User{Uid: loginID}
	err := model.GetOne(context.TODO(), &dbinfo)
	if err != nil {
		return "", err
	}
	if dbinfo.Passwd != loginPwd {
		return "", errors.New("wrong username or password")
	} else {
		return token.Newtoken(loginID)
	}
}
