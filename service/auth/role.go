package auth

import (
	"context"
	"main/dao"
)

func SetRole(ctx context.Context, ccnuid string, role int) error {
	info := &dao.User{CCNUid: ccnuid}
	err := info.FindByCCNUid(ctx)
	if err != nil {
		return err
	}
	info.Role = role
	err = info.Update(ctx)
	if err != nil {
		return err
	}
	return nil
}
