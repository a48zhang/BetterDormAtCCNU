package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/service/auth"
)

type LoginRequest struct {
	CCNUid string `json:"ccnuid" bson:"ccnuid"`
	Name   string `json:"name" bson:"name"`
	Passwd string `json:"passwd" bson:"passwd"`
}

// Login godoc
//
// @Summary 登录
// @Tags auth
// @Accept	json
// @Produce json
// @Param Login body LoginRequest true "登录信息，只需要填写Name和Passwd"
// @Success	200	{object} Resp
// @Failure	400	{object} Resp
// @Router  /login [post]
func Login(ctx *gin.Context) {
	req := LoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseError(ctx, 400, "Bad Request")
		return
	}
	token, err := auth.Login(ctx.Request.Context(), &model.User{Name: req.Name, Passwd: req.Passwd})
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	ResponseOK(ctx, gin.H{"token": token})
}

type RegisterRequest struct {
	CCNUid string `json:"ccnuid" bson:"ccnuid"`
	Name   string `json:"name" bson:"name"`
	Passwd string `json:"passwd" bson:"passwd"`
}

// Register godoc
//
// @Summary 注册
// @Tags auth
// @Accept	json
// @Produce json
// @Param Register body RegisterRequest true "注册信息"
// @Success	200	{object} Resp
// @Failure	400	{object} Resp
// @Router  /register [post]
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
	ResponseOK(ctx, nil)
}
