package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type structs interface {
	User
	Col() *mongo.Collection
}

func GetOne[T structs](ctx context.Context, info *T) error {
	return (*info).Col().FindOne(ctx, *info).Decode(&info)
}

func GetMany[T structs](ctx context.Context, info *T) ([]T, error) {
	cursor, err := (*info).Col().Find(ctx, *info)
	if err != nil {
		return nil, err
	}
	var result []T
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Insert[T structs](ctx context.Context, info *T) error {
	_, err := (*info).Col().InsertOne(ctx, *info)
	return err
}

func Delete[T structs](ctx context.Context, info *T) error {
	_, err := (*info).Col().DeleteOne(ctx, *info)
	return err
}

func Update[T structs](ctx context.Context, info *T) error {
	// TODO
}
