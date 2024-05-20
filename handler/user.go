package handler

import (
	"github.com/gin-gonic/gin"
	"main/dao"
	"main/pkg"
	"main/service/auth"
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
//	@Summary		修改当前登录用户信息
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

// UpdateUserInfo godoc
//
//	@Summary		修改指定用户信息
//	@Description	Update user's info
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			token	header		string		true	"token"
//	@Param			info	body		dao.User	true	"info"
//	@Param			ccnuid	query		string	true	"ccnuid"
//	@Success		200		{object}	Resp
//	@Failure		500		{object}	Resp
//	@Router			/users/update [post]
func UpdateUserInfo(ctx *gin.Context) {
	_role := ctx.GetString("role")
	ccnuid := ctx.Query("ccnuid")
	role, err := strconv.Atoi(_role)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	if role < 2 {
		ResponseError(ctx, 403, "role not satisfied")
		return
	}
	info := dao.User{CCNUid: ccnuid}
	err = info.FindByCCNUid(ctx)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	err = ctx.BindJSON(&info)
	err = userservice.UpdateUserInfo(ctx.Request.Context(), info)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, info)
}

type SetUserRoleRequest struct {
	CCNUid string `json:"ccnuid" bson:"ccnuid"`
	Role   int    `json:"role" bson:"role"`
}

// SetUserRole godoc
//
//	@Summary	设置用户角色
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		SetUserRole	body		SetUserRoleRequest	true	"设置用户角色"
//	@Success	200			{object}	Resp
//	@Failure	400			{object}	Resp
//	@Router		/users/setrole [post]
func SetUserRole(ctx *gin.Context) {
	req := SetUserRoleRequest{}
	r := ctx.GetString("role")
	if r == "0" {
		ResponseError(ctx, 403, "role not satisfied")
		return
	}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	err = auth.SetRole(ctx.Request.Context(), req.CCNUid, req.Role)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, nil)
}
