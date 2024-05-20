package dao

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
)

type Password struct {
	MD5Hash    string
	SHA256Hash string
	Salt       string
}

func (p *Password) Is(pwd string) bool {
	pwd += p.Salt
	h1 := sha256.New()
	_, _ = io.WriteString(h1, pwd)
	sha2h := h1.Sum(nil)
	h2 := md5.New()
	_, _ = io.WriteString(h2, pwd)
	md5h := h2.Sum(nil)
	return p.SHA256Hash == string(sha2h) && p.MD5Hash == string(md5h)
}

type User struct {
	MID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Uid    string             `json:"uid" bson:"uid"`
	CCNUid string             `json:"ccnuid" bson:"ccnuid"`
	Name   string             `json:"name" bson:"name"`
	Passwd Password           `json:"passwd" bson:"passwd"`
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
	return remove(ctx, &u)
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
