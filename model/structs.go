package model

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	Uid    string `json:"uid" bson:"uid"`
	CCNUid string `json:"ccnuid" bson:"ccnuid"`
	Name   string `json:"name" bson:"name"`
	Passwd string `json:"passwd" bson:"passwd"`
	Role   int    `json:"role" bson:"role"` // role: 0 for student, 1 for teacher, 2 for school
	Valid  int    `json:"valid" bson:"valid"`
}

func (u User) Col() *mongo.Collection {
	return DB.Collection("user")
}

func (u User) ID() (string, string) {
	return "ccnuid", u.CCNUid
}

type Form struct {
	Fid       string `json:"fid" bson:"fid"`
	StudentID string `json:"student_id" bson:"student_id"`
	TeacherID string `json:"teacher_id" bson:"teacher_id"`
	CreateAt  string `json:"create_at" bson:"create_at"`
	FromDorm  string `json:"from_dorm" bson:"from_dorm"`
	ToDorm    string `json:"to_dorm" bson:"to_dorm"`
	Context   string `json:"context" bson:"context"`
	Status    int    `json:"status" bson:"status"`
}

func (f Form) Col() *mongo.Collection {
	return DB.Collection("form")
}

func (f Form) ID() (string, string) {
	return "fid", f.Fid
}
