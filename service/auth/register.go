package auth

import (
	"context"
	"main/dao"
)

func Register(ctx context.Context, info *dao.User) error {
	return info.Insert(ctx)
}
