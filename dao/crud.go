package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Elem interface {
	Col() *mongo.Collection
	ID() primitive.ObjectID
}

type DAO interface {
	Insert(context.Context) error
	Update(context.Context) error
	Delete(context.Context) error
	Find(context.Context) error
	FindAll(context.Context) ([]any, error)
}

func find[T Elem](ctx context.Context, info *T) error {
	return (*info).Col().FindOne(ctx, bson.D{{"_id", (*info).ID()}}).Decode(&info)
}

func findAll[T Elem](ctx context.Context, info *T) ([]T, error) {
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

func insert[T Elem](ctx context.Context, info *T) error {
	_, err := (*info).Col().InsertOne(ctx, *info)
	return err
}

func remove[T Elem](ctx context.Context, info *T) error {
	_, err := (*info).Col().DeleteOne(ctx, *info)
	return err
}

func update[T Elem](ctx context.Context, info *T) error {
	_, err := (*info).Col().ReplaceOne(ctx, bson.D{{"_id", (*info).ID()}}, *info)
	return err
}

func Insert[T Elem](ctx context.Context, info *T) error {
	return insert(ctx, info)
}
func Update[T Elem](ctx context.Context, info *T) error {
	return update(ctx, info)
}
func Delete[T Elem](ctx context.Context, info *T) error {
	return remove(ctx, info)
}
func Find[T Elem](ctx context.Context, info *T) error {
	return find(ctx, info)
}
func FindAll[T Elem](ctx context.Context, info *T) ([]T, error) {
	return findAll(ctx, info)
}
