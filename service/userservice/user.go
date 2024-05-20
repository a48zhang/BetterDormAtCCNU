package userservice

import (
	"context"
	"main/dao"
)

func GetUserInfo(ctx context.Context, uid string) (dao.User, error) {
	info := &dao.User{CCNUid: uid}
	err := info.FindByCCNUid(ctx)
	info.Passwd = dao.Password{}
	return *info, err
}

func UpdateUserInfo(ctx context.Context, info dao.User) error {
	req := &dao.User{}
	*req = info
	err := req.Update(ctx)
	return err
}
