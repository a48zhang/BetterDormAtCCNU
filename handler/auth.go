package handler

import (
	"github.com/gin-gonic/gin"
	"main/dao"
	"main/service/auth"
)

type LoginRequest struct {
	CCNUid string `json:"ccnuid" bson:"ccnuid"`
	Name   string `json:"name" bson:"name"`
	Passwd string `json:"passwd" bson:"passwd"`
}

// Login godoc
//
//	@Summary	登录
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		Login	body		LoginRequest	true	"登录信息，只需要填写Name和Passwd"
//	@Success	200		{object}	Resp
//	@Failure	400		{object}	Resp
//	@Router		/login [post]
func Login(ctx *gin.Context) {
	req := LoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	token, err := auth.Login(ctx.Request.Context(), &dao.User{Name: req.Name, Passwd: req.Passwd})
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
//	@Summary	注册
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		Register	body		RegisterRequest	true	"注册信息"
//	@Success	200			{object}	Resp
//	@Failure	400			{object}	Resp
//	@Router		/register [post]
func Register(ctx *gin.Context) {
	req := RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	err = auth.Register(ctx.Request.Context(), &dao.User{CCNUid: req.CCNUid, Name: req.Name, Passwd: req.Passwd})
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	ResponseOK(ctx, nil)
}
