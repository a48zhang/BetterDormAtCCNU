package userservice

import (
	"context"
	"main/model"
)

func GetUserInfo(ctx context.Context, uid string) (model.User, error) {
	info := model.User{CCNUid: uid}
	err := model.GetOne(ctx, &info)
	return info, err
}

func UpdateUserInfo(ctx context.Context, info model.User) error {
	err := model.Update(ctx, &info)
	return err
}
