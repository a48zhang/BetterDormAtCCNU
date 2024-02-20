package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/service/auth"
)

type LoginRequest struct {
	CCNUid string `json:"ccnuid" bson:"ccnuid"`
	Passwd string `json:"passwd"`
}

func Login(ctx *gin.Context) {
	req := LoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseError(ctx, 400, "Bad Request")
		return
	}
	token, err := auth.Login(ctx.Request.Context(), &model.User{CCNUid: req.CCNUid, Passwd: req.Passwd})
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	ResponseOK(ctx, gin.H{"token": token})
}

type RegisterRequest struct {
	CCNUid string `json:"ccnuid" bson:"ccnuid"`
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

func Register(ctx *gin.Context) {
	req := RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseError(ctx, 400, "Bad Request")
		return
	}
	err = auth.Register(ctx.Request.Context(), &model.User{CCNUid: req.CCNUid, Name: req.Name, Passwd: req.Passwd})
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	ResponseOK(ctx, gin.H{"warning": "this api will be deprecated in the future."})
}
