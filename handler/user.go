package handler

import (
	"github.com/gin-gonic/gin"
	"main/service/userservice"
)

// GetUserInfo godoc
// @Summary Get user info
// @Description Get user info
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} Resp
// @Failure 500 {object} Resp
// @Router /users [get]
func GetUserInfo(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	info, err := userservice.GetUserInfo(ctx.Request.Context(), uid)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, info)
}

// UpdateUserInfo godoc
// @Summary Update user info
// @Description Update user info
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body model.User true "info"
// @Param uid path string true "uid"
// @Success 200 {object} Resp
// @Failure 500 {object} Resp
// @Router /users [post]
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
