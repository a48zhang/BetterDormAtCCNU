package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/model"
	"main/pkg/token"
	"net/http"
)

func TokenParser(r *gin.Context) {
	tokenString := r.GetHeader("token")
	if tokenString == "" {
		r.AbortWithError(http.StatusForbidden, errors.New("token not found"))
		return
	}
	parsetoken, claims, err := token.Parsetoken(tokenString)
	if err != nil || !parsetoken.Valid {
		r.AbortWithError(http.StatusForbidden, errors.New("token not valid"))
		return
	}
	r.Set("uid", claims.UID)
	r.Set("expiresAt", claims.ExpiresAt)
	info := &model.User{CCNUid: claims.UID}
	err = info.FindByCCNUid(r)
	if err != nil {
		r.AbortWithError(http.StatusForbidden, errors.New("user not valid"))
	}
	r.Set("Role", info.Role)
	r.Next()
}
