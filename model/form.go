package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Form struct {
	MID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Fid           string             `json:"fid" bson:"fid"`
	StudentID     string             `json:"student_id" bson:"student_id"`
	College       string             `json:"college" bson:"college"`
	Contact       string             `json:"contact" bson:"contact"`
	TeacherID     string             `json:"teacher_id" bson:"teacher_id"`
	CreateAt      string             `json:"create_at" bson:"create_at"`
	FromDorm      string             `json:"from_dorm" bson:"from_dorm"`
	ToDorm        string             `json:"to_dorm" bson:"to_dorm"`
	ToBed         string             `json:"to_bed" bson:"to_bed"`
	Context       string             `json:"context" bson:"context"`
	TeacherAdvice string             `json:"teacher_advice" bson:"teacher_advice"`
	Status        int                `json:"status" bson:"status"`
}

func (f *Form) Insert(ctx context.Context) error {
	return insert(ctx, &f)
}

func (f *Form) Update(ctx context.Context) error {
	return update(ctx, &f)
}

func (f *Form) Delete(ctx context.Context) error {
	return delete(ctx, &f)
}

func (f *Form) Find(ctx context.Context) error {
	return find(ctx, &f)
}

func (f *Form) FindAll(ctx context.Context) ([]any, error) {
	data, err := findAll(ctx, &f)
	if err != nil {
		return nil, err
	}
	ret := make([]any, len(data))
	for i, v := range data {
		ret[i] = v
	}
	return ret, nil
}

func (f *Form) FindByStudentId(ctx context.Context, sid string) ([]any, error) {
	cursor, err := f.Col().Find(ctx, bson.M{"student_id": sid})
	if err != nil {
		return nil, err
	}
	var result []Form
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

func (f *Form) Col() *mongo.Collection {
	return DB.Collection("form")
}

func (f *Form) ID() primitive.ObjectID {
	return f.MID
}
