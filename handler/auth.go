package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/service/auth"
)

type LoginRequest struct {
	Uid    string `json:"uid"`
	Passwd string `json:"passwd"`
}

func Login(ctx *gin.Context) {
	req := LoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseError(ctx, 400, "Bad Request")
		return
	}
	token, err := auth.Login(&model.User{Uid: req.Uid, Passwd: req.Passwd})
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	ResponseOK(ctx, gin.H{"token": token})
}

type RegisterRequest struct {
	Uid    string `json:"uid"`
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
	err = auth.Register(&model.User{Uid: req.Uid, Name: req.Name, Passwd: req.Passwd})
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	ResponseOK(ctx, gin.H{"warning": "this api will be deprecated in the future."})
}
