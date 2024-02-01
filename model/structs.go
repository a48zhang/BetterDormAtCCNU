package model

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	Uid    string `json:"uid" bson:"uid"`
	Name   string `json:"name" bson:"name"`
	Passwd string `json:"passwd" bson:"passwd"`
	Role   int    `json:"role" bson:"role"`
}

func (u User) Col() *mongo.Collection {
	return DB.Collection("user")
}

type Form struct {
	Fid          string `json:"fid" bson:"fid"`
	StudentID    string `json:"student_id" bson:"student_id"`
	TeacherID    string `json:"teacher_id" bson:"teacher_id"`
	CreateAt     string `json:"create_at" bson:"create_at"`
	LastModifyAt string `json:"last_modify_at" bson:"last_modify_at"`
	FromDorm     string `json:"from_dorm" bson:"from_dorm"`
	ToDorm       string `json:"to_dorm" bson:"to_dorm"`
	Context      string `json:"context" bson:"context"`
	Status       int    `json:"status" bson:"status"`
}
