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
	TeacherID string `json:"teacher_id" bson:"teacher_id"`
	FromDorm  string `json:"from_dorm" bson:"from_dorm"`
	ToDorm    string `json:"to_dorm" bson:"to_dorm"`
	Context   string `json:"context" bson:"context"`
}

func CreateForm(ctx *gin.Context) {
	info := FormRequest{}
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	err = form.CreateForm(ctx.Request.Context(), model.Form{
		StudentID: info.StudentID,
		TeacherID: info.TeacherID,
		CreateAt:  time.Now().String(),
		FromDorm:  info.FromDorm,
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

func GetForms(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	forms, err := form.GetUsersForm(ctx.Request.Context(), uid)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, forms)
}

func GetOneForm(ctx *gin.Context) {
	fid := ctx.Query("fid")
	data := model.Form{Fid: fid}
	err := model.GetOne(ctx.Request.Context(), &data)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, data)
}

type CheckFormRequest struct {
	Fid string `json:"fid" bson:"fid"`
	Opt int    `json:"opt" bson:"opt"`
}

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
	err = form.CheckForm(ctx.Request.Context(), info.Fid, info.Opt)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, nil)
}
