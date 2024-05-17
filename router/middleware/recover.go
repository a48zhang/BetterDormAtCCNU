package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Recovery(r *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			r.AbortWithStatusJSON(500, gin.H{
				"code": 500,
				"data": gin.H{
					"request_id": r.GetString("request_id"),
					"time":       time.Now().Format("2006-01-02 15:04:05"),
				},
				"message": "Internal Server Error",
			})
			ServiceErrorCount++
			log.Println(r.Request, r.Keys, r.Errors)
		}
	}()
	r.Next()
}
