package handler

import "github.com/gin-gonic/gin"

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Response(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func ResponseOK(ctx *gin.Context, data interface{}) {
	Response(ctx, 200, data, "OK")
}

func ResponseError(ctx *gin.Context, code int, msg string) {
	Response(ctx, code, nil, msg)
}
