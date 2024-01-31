package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/pkg/token"
	"net/http"
)

func TokenParser(r *gin.Context) {
	tokenString := r.GetHeader("Authorization")
	if tokenString == "" {
		r.AbortWithError(http.StatusForbidden, errors.New("token not found"))
		return
	}
	token, claims, err := token.Parsetoken(tokenString)
	if err != nil || !token.Valid {
		r.AbortWithError(http.StatusForbidden, errors.New("token not valid"))
		return
	}
	r.Set("userID", claims.UID)
	r.Set("expiresAt", claims.ExpiresAt)
	r.Next()
}
