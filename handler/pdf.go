package handler

import (
	"github.com/gin-gonic/gin"
	"main/service/form"
	"main/service/pdfgen"
	"main/service/userservice"
)

type PDFRequest struct {
	Name    string `json:"name"`
	School  string `json:"school"`
	Sid     string `json:"sid"`
	Contact string `json:"contact"`
	From    string `json:"from"`
	To      string `json:"to"`
	Text0   string `json:"text0"`
	Text1   string `json:"text1"`
	Text2   string `json:"text2"`
	Text3   string `json:"text3"`
	Text4   string `json:"text4"`
}

// GenFormPDF godoc
//
//	@Summary		根据上传数据生成表单PDF
//	@Description	Generate PDF for form
//	@Tags			form
//	@Accept			json
//	@Produce		application/pdf
//	@Param			token	header		string		true	"token"
//	@Param			info	body		PDFRequest	true	"info"
//	@Success		200		{file}		application/pdf
//	@Failure		400		{object}	Resp
//	@Failure		500		{object}	Resp
//	@Router			/forms/pdf [post]
func GenFormPDF(ctx *gin.Context) {
	info := pdfgen.PdfForm{}
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		ResponseError(ctx, 400, err.Error())
	}

	id, err := pdfgen.GenPDFForForm(ctx, info)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
	}
	ctx.File(pdfgen.GetFilePath(id))
	pdfgen.Waitlist[id] = 5
}

// GenFormByID godoc
//
//	@Summary		根据表单ID生成表单PDF
//	@Description	Generate PDF for form by ID
//	@Tags			form
//	@Accept			json
//	@Produce		application/pdf
//	@Param			token	header		string	true	"token"
//	@Param			id		query		string	true	"form id"
//	@Success		200		{file}		application/pdf
//	@Failure		400		{object}	Resp
//	@Failure		500		{object}	Resp
//	@Router			/forms/pdf [get]
func GenFormByID(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	id := ctx.Query("id")
	getForm, err := form.GetOneForm(ctx, id)
	if err != nil {
		return
	}
	userInfo, err := userservice.GetUserInfo(ctx, uid)
	if err != nil {
		return
	}

	info := pdfgen.PdfForm{
		ID:      "",
		Name:    userInfo.Name,
		School:  getForm.College,
		Sid:     getForm.StudentID,
		Contact: getForm.Contact,
		From:    getForm.FromDorm + " " + getForm.FromBed,
		To:      getForm.ToDorm + " " + getForm.ToBed,
		Text0:   getForm.Context,
		Text1:   getForm.TeacherAdvice,
		Text2:   getForm.CommunityAdvice,
		Text3:   getForm.XGBAdvice,
		Text4:   getForm.HQBZBAdvice,
	}

	id, err = pdfgen.GenPDFForForm(ctx, info)
	if err != nil {
		ResponseError(ctx, 500, err.Error())
	}
	ctx.File(pdfgen.GetFilePath(id))
	pdfgen.Waitlist[id] = 5
}

func GenPDFCallback(ctx *gin.Context) {
	id := ctx.Query("id")
	pdfgen.Waitlist[id] = 2
	Response(ctx, 200, id, id)
}
