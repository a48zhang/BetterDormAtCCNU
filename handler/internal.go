package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/dao"
	"main/pkg/token"
	"time"
)

// TokenStatus godoc
//
//	@Summary		查看token状态
//	@Description	Check token status
//	@Tags			auth
//	@Produce		json
//	@Param			token	header		string	true	"token"
//	@Success		200		{object}	Resp
//	@Failure		403		{object}	Resp
func TokenStatus(ctx *gin.Context) {
	tokenString := ctx.GetHeader("token")
	parsetoken, claims, err := token.Parsetoken(tokenString)
	if err != nil {
		ResponseError(ctx, 403, "failed to parse token")
		return
	}
	ResponseOK(ctx, gin.H{
		"token": tokenString,
		"claims": gin.H{
			"UserMID":     claims.UserMID,
			"Username":    claims.Username,
			"UID(CCNUID)": claims.UID,
			"IssuedAt":    time.Unix(claims.IssuedAt, 0).String(),
			"ExpiresAt":   time.Unix(claims.ExpiresAt, 0).String(),
			"Issuer":      claims.Issuer,
		},
		"valid": parsetoken.Valid,
	})
}

// TokenRefresh godoc
//
//	@Summary		刷新token
//	@Description	Refresh token
//	@Tags			auth
//	@Produce		json
//	@Param			token	header		string	true	"token"
//	@Success		200		{object}	Resp
//	@Failure		403		{object}	Resp
func TokenRefresh(ctx *gin.Context) {
	MID, _ := ctx.Get("UserMID")
	info := dao.User{MID: MID.(primitive.ObjectID)}
	err := info.Find(ctx)
	if err != nil {
		ResponseError(ctx, 404, "user not found")
		return
	}
	newtoken, _ := token.NewTokenForUser(info)
	ResponseOK(ctx, gin.H{
		"token": newtoken,
	})
}
