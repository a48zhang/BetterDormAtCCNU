package handler

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Uid    string `json:"uid"`
	Passwd string `json:"passwd"`
}

func StudentLogin(ctx *gin.Context) {
	req := LoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ResponseError(ctx, 400, "Bad Request")
		return
	}

}
