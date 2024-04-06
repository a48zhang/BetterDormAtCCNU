package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Report struct {
	RID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Uid     string             `json:"uid" bson:"uid"`
	Info    string             `json:"info" bson:"info"`
	Contact string             `json:"contact" bson:"contact"`
	Comment string             `json:"comment" bson:"comment"`
}

func (r *Report) Insert(ctx context.Context) error {
	return insert(ctx, &r)
}

func (r *Report) Update(ctx context.Context) error {
	return update(ctx, &r)
}

func (r *Report) Delete(ctx context.Context) error {
	return delete(ctx, &r)
}

func (r *Report) Find(ctx context.Context) error {
	return find(ctx, &r)
}

func (r *Report) FindAll(ctx context.Context) ([]any, error) {
	data, err := findAll(ctx, &r)
	if err != nil {
		return nil, err
	}
	ret := make([]any, len(data))
	for i, v := range data {
		ret[i] = v
	}
	return ret, nil
}

func (r *Report) FindByUid(ctx context.Context) ([]any, error) {
	cursor, err := r.Col().Find(ctx, bson.M{"uid": r.Uid})
	if err != nil {
		return nil, err
	}
	var result []Report
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	res := make([]any, len(result))
	for i, v := range result {
		res[i] = v
	}
	return res, nil
}

func (r *Report) Col() *mongo.Collection {
	return DB.Collection("report")
}

func (r *Report) ID() primitive.ObjectID {
	return r.RID
}
