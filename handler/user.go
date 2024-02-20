package handler

import (
	"github.com/gin-gonic/gin"
	"main/service/userservice"
)

func GetUserInfo(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	info, err := userservice.GetUserInfo(ctx.Request.Context(), uid)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, info)
}

func UpdateUserInfo(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	info, err := userservice.GetUserInfo(ctx.Request.Context(), uid)
	err = ctx.BindJSON(&info)
	err = userservice.UpdateUserInfo(ctx.Request.Context(), info)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, info)
}
