package handler

import (
	"github.com/gin-gonic/gin"
	"main/dao"
	"main/pkg"
	"main/service/userservice"
	"strconv"
)

// GetUserInfo godoc
//
//	@Summary		获取用户信息
//	@Description	Get user info
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			token	header		string	true	"token"
//	@Success		200		{object}	Resp
//	@Failure		500		{object}	Resp
//	@Router			/users [get]
func GetUserInfo(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	info, err := userservice.GetUserInfo(ctx.Request.Context(), uid)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	role, _ := pkg.Get5thBit(uid)
	Response(ctx, 200, info, role)
}

// UpdateMyInfo godoc
//
//	@Summary		修改用户信息
//	@Description	Update user's info
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			token	header		string		true	"token"
//	@Param			info	body		dao.User	true	"info"
//	@Success		200		{object}	Resp
//	@Failure		500		{object}	Resp
//	@Router			/users [post]
func UpdateMyInfo(ctx *gin.Context) {
	uid := ctx.GetString("uid")

	info := dao.User{CCNUid: uid}
	info.FindByCCNUid(ctx)
	err := ctx.BindJSON(&info)
	err = userservice.UpdateUserInfo(ctx.Request.Context(), info)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, info)
}

func UpdateUserInfo(ctx *gin.Context) {
	_role := ctx.GetString("role")
	role, err := strconv.Atoi(_role)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	if role < 2 {
		ResponseError(ctx, 403, "role not satisfied")
		return
	}
}
