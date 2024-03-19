package userservice

import (
	"context"
	"main/model"
)

func GetUserInfo(ctx context.Context, uid string) (model.User, error) {
	info := &model.User{CCNUid: uid}
	err := info.FindByCCNUid(ctx)
	return *info, err
}

func UpdateUserInfo(ctx context.Context, info model.User) error {
	req := &model.User{}
	*req = info
	err := req.FindByCCNUid(ctx)
	if err != nil {
		return err
	}
	err = req.Update(ctx)
	return err
}
