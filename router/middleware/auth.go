package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"main/dao"
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
	r.Set("UserMID", claims.UserMID)
	r.Set("IssuedAt", claims.IssuedAt)
	r.Set("expiresAt", claims.ExpiresAt)
	// Note: context[uid] = User.CCNUID
	info := &dao.User{MID: claims.UserMID}
	err = info.Find(r)
	if err != nil {
		log.Println(*claims, err)
		r.AbortWithError(http.StatusForbidden, errors.New("user not valid"))
	}
	r.Set("uid", info.CCNUid)
	r.Set("Username", info.Name)
	r.Set("Role", info.Role)
	r.Next()
}
