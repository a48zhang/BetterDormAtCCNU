package auth

import (
	"context"
	"main/model"
)

func Register(ctx context.Context, info *model.User) error {
	return model.Insert(ctx, info)
}
