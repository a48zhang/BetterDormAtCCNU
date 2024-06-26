package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/dao"
	"main/service/form"
	"math"
	"time"
)

type FormRequest struct {
	MID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID       string             `json:"student_id" bson:"student_id"`
	College         string             `json:"college" bson:"college"`
	Contact         string             `json:"contact" bson:"contact"`
	TeacherID       string             `json:"teacher_id" bson:"teacher_id"`
	CreateAt        string             `json:"create_at" bson:"create_at"`
	FromDorm        string             `json:"from_dorm" bson:"from_dorm"`
	FromBed         string             `json:"from_bed" bson:"from_bed"`
	ToDorm          string             `json:"to_dorm" bson:"to_dorm"`
	ToBed           string             `json:"to_bed" bson:"to_bed"`
	Context         string             `json:"context" bson:"context"`
	TeacherAdvice   string             `json:"teacher_advice" bson:"teacher_advice"`
	CommunityAdvice string             `json:"community_advice" bson:"community_advice"`
	XGBAdvice       string             `json:"xgb_advice" bson:"xgb_advice"`
	HQBZBAdvice     string             `json:"hqzb_advice" bson:"hqzb_advice"`
	Status          int                `json:"status" bson:"status"`
}

// CreateForm godoc
//
//	@Summary	提交新换宿申请表格
//	@Tags		form
//	@Accept		json
//	@Produce	json
//	@Param		token	header		string		true	"token"
//	@Param		Form	body		FormRequest	true	"换宿申请表格"
//	@Success	200		{object}	Resp
//	@Failure	400		{object}	Resp
//	@Router		/forms/create [post]
func CreateForm(ctx *gin.Context) {
	info := FormRequest{}
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
		return
	}
	err = form.CreateForm(ctx.Request.Context(), dao.Form{
		StudentID: info.StudentID,
		College:   info.College,
		TeacherID: info.TeacherID,
		Contact:   info.Contact,
		CreateAt:  time.Now().String(),
		FromDorm:  info.FromDorm,
		FromBed:   info.FromBed,
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
//	@Summary	获取用户的换宿申请表格
//	@Tags		form
//	@Accept		json
//	@Produce	json
//	@Param		token	header		string	true	"token"
//	@Success	200		{object}	Resp
//	@Failure	500		{object}	Resp
//	@Router		/forms/my [get]
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
//	@Summary	获取单个换宿申请表格
//	@Tags		form
//	@Accept		json
//	@Produce	json
//	@Param		token	header		string	true	"token"
//	@Param		fid		query		string	true	"fid"
//	@Success	200		{object}	Resp
//	@Failure	500		{object}	Resp
//	@Router		/forms [get]
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
//	@Summary		审核换宿申请表格
//	@Description	opt: 0 for waiting, 1 for teacher approve, -1 for teacher reject, 2 for school approve, -2 for school reject
//	@Description	if role < abs(opt), permission denied
//	@Tags			form
//	@Accept			json
//	@Produce		json
//	@Param			token		header		string				true	"token"
//	@Param			CheckForm	body		CheckFormRequest	true	"审核信息"
//	@Success		200			{object}	Resp
//	@Failure		400			{object}	Resp
//	@Failure		403			{object}	Resp
//	@Failure		500			{object}	Resp
//	@Router			/forms/check [post]
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

// GetAssignedForms godoc
//
//	@Summary	获取分配给指定角色的换宿申请表格
//	@Tags		form
//	@Accept		json
//	@Produce	json
//	@Param		token	header		string	true	"token"
//	@Success	200		{object}	Resp
//	@Failure	500		{object}	Resp
//	@Router		/forms/assigned [get]
func GetAssignedForms(ctx *gin.Context) {
	role := ctx.GetInt("role")
	uid := ctx.GetString("uid")
	forms, err := form.GetAssignedForms(ctx.Request.Context(), uid, role)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
		return
	}
	ResponseOK(ctx, forms)
}
