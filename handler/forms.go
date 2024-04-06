package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/service/form"
	"math"
	"time"
)

type FormRequest struct {
	StudentID string `json:"student_id" bson:"student_id"`
	College   string `json:"college" bson:"college"`
	Contact   string `json:"contact" bson:"contact"`
	TeacherID string `json:"teacher_id" bson:"teacher_id"`
	FromDorm  string `json:"from_dorm" bson:"from_dorm"`
	ToDorm    string `json:"to_dorm" bson:"to_dorm"`
	ToBed     string `json:"to_bed" bson:"to_bed"`
	Context   string `json:"context" bson:"context"`
}

// CreateForm godoc
//
// @Summary 提交新换宿申请表格
// @Tags form
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Param Form body FormRequest true "换宿申请表格"
// @Success	200	{object} Resp
// @Failure	400	{object} Resp
// @Router  /forms/create [post]
func CreateForm(ctx *gin.Context) {
	info := FormRequest{}
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	err = form.CreateForm(ctx.Request.Context(), model.Form{
		StudentID: info.StudentID,
		College:   info.College,
		TeacherID: info.TeacherID,
		Contact:   info.Contact,
		CreateAt:  time.Now().String(),
		FromDorm:  info.FromDorm,
		ToBed:     info.ToBed,
		ToDorm:    info.ToDorm,
		Context:   info.Context,
		Status:    0,
	})
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, nil)
}

// GetForms godoc
//
// @Summary 获取用户的换宿申请表格
// @Tags form
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Success	200	{object} Resp
// @Failure	500	{object} Resp
// @Router  /forms/my [get]
func GetForms(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	forms, err := form.GetUsersForm(ctx.Request.Context(), uid)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, forms)
}

// GetOneForm godoc
//
// @Summary 获取单个换宿申请表格
// @Tags form
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Param fid query string true "fid"
// @Success	200	{object} Resp
// @Failure	500	{object} Resp
// @Router  /forms [get]
func GetOneForm(ctx *gin.Context) {
	fid := ctx.Query("fid")
	data, err := form.GetOneForm(ctx.Request.Context(), fid)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, data)
}

type CheckFormRequest struct {
	Fid    string `json:"fid" bson:"fid"`
	Opt    int    `json:"opt" bson:"opt"`
	Advice string `json:"advice" bson:"advice"`
}

// CheckForm godoc
//
// @Summary 审核换宿申请表格
// @Tags form
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Param CheckForm body CheckFormRequest true "审核信息"
// @Success	200	{object} Resp
// @Failure	400	{object} Resp
// @Failure	403	{object} Resp
// @Failure	500	{object} Resp
// @Router  /forms/check [post]
func CheckForm(ctx *gin.Context) {
	role := ctx.GetInt("role")
	info := CheckFormRequest{}
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	if int(math.Abs(float64(info.Opt))) > role {
		ResponseError(ctx, 403, "Permission denied")
		return
	}
	err = form.CheckForm(ctx.Request.Context(), info.Fid, info.Opt, info.Advice)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, nil)
}
