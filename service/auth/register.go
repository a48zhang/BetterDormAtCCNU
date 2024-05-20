package auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/dao"
)

func Register(ctx context.Context, info *dao.User) error {
	// check if user exists
	UserExists := errors.New("user exists")
	f := &dao.User{
		Name: info.Name,
	}
	err := f.FindByName(ctx)
	if err == nil {
		return UserExists
	}
	f = &dao.User{
		CCNUid: info.CCNUid,
	}
	err = f.FindByCCNUid(ctx)
	if err == nil {
		return UserExists
	}

	info.MID = primitive.NilObjectID
	return info.Insert(ctx)
}
