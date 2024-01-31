package model

type User struct {
	Uid    string `json:"uid" bson:"uid"`
	Name   string `json:"name" bson:"name"`
	Passwd string `json:"passwd" bson:"passwd"`
}
