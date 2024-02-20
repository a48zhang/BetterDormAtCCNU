package auth

import (
	"context"
	"errors"
	"main/model"
	"main/pkg/token"
)

func Login(ctx context.Context, LoginInfo *model.User) (string, error) {
	loginID := LoginInfo.CCNUid
	loginPwd := LoginInfo.Passwd
	dbinfo := model.User{CCNUid: loginID}
	err := model.GetOne(ctx, &dbinfo)
	if err != nil {
		return "", err
	}
	if dbinfo.Passwd != loginPwd {
		return "", errors.New("wrong username or password")
	} else {
		return token.Newtoken(loginID)
	}
}
