package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogInfo struct {
	MID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RequestID string             `json:"request_id" bson:"request_id"`
	Data      gin.H              `json:"data" bson:"data"`
}

func (l *LogInfo) Col() *mongo.Collection {
	return DB.Collection("bd_logs")
}

func (l *LogInfo) ID() primitive.ObjectID {
	return l.MID
}

func (l *LogInfo) Insert(ctx context.Context) error {
	return insert(ctx, &l)
}

func (l *LogInfo) Update(ctx context.Context) error {
	return update(ctx, &l)
}

func (l *LogInfo) Delete(ctx context.Context) error {
	return remove(ctx, &l)
}

func (l *LogInfo) Find(ctx context.Context) error {
	return find(ctx, &l)
}
