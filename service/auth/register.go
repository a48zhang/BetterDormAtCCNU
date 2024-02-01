package auth

import (
	"context"
	"main/model"
)

func Register(info *model.User) error {
	return model.Insert(context.TODO(), info)
}
