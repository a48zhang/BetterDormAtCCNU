package handler

import (
	"github.com/gin-gonic/gin"
	"main/service/userservice"
)

type ReportRequest struct {
	Info    string `json:"info" bson:"info"`
	Contact string `json:"contact" bson:"contact"`
}

// NewReport godoc
//
// @Summary 提交新反馈
// @Tags user
// @Accept	json
// @Param token header string true "token"
// @Produce json
// @Param Report body ReportRequest true "反馈信息"
// @Success	200	{object} Resp
// @Failure	400	{object} Resp
// @Failure	500	{object} Resp
// @Router  /report [post]
func NewReport(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	info := ReportRequest{}
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	err = userservice.NewReport(ctx, uid, info.Info, info.Contact)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, nil)
}

// MyReport godoc
//
// @Summary 我的反馈
// @Tags user
// @Param token header string true "token"
// @Produce json
// @Success	200	{object} Resp
// @Failure	500	{object} Resp
// @Router  /report [get]
func MyReport(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	data, err := userservice.MyReport(ctx, uid)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, data)
}
