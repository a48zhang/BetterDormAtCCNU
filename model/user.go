package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	MID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Uid    string             `json:"uid" bson:"uid"`
	CCNUid string             `json:"ccnuid" bson:"ccnuid"`
	Name   string             `json:"name" bson:"name"`
	Passwd string             `json:"passwd" bson:"passwd"`
	Role   int                `json:"role" bson:"role"` // role: 0 for student, 1 for teacher, 2 for school
	School string             `json:"school" bson:"school"`
	Stage  string             `json:"stage" bson:"stage"`
	Valid  int                `json:"valid" bson:"valid"`
}

func (u *User) Insert(ctx context.Context) error {
	return insert(ctx, &u)
}

func (u *User) Update(ctx context.Context) error {
	return update(ctx, &u)
}

func (u *User) Delete(ctx context.Context) error {
	return delete(ctx, &u)
}

func (u *User) Find(ctx context.Context) error {
	return find(ctx, &u)
}

func (u *User) FindByName(ctx context.Context) error {
	return u.Col().FindOne(ctx, bson.D{{"name", u.Name}}).Decode(u)
}

func (u *User) FindByCCNUid(ctx context.Context) error {
	return u.Col().FindOne(ctx, bson.D{{"ccnuid", u.CCNUid}}).Decode(u)
}

func (u *User) FindAll(ctx context.Context) ([]any, error) {
	data, err := findAll(ctx, &u)
	if err != nil {
		return nil, err
	}
	ret := make([]any, len(data))
	for i, v := range data {
		ret[i] = v
	}
	return ret, nil
}

func (u *User) Col() *mongo.Collection {
	return DB.Collection("user")
}

func (u *User) ID() primitive.ObjectID {
	return u.MID
}
