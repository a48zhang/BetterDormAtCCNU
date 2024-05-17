package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"main/pkg/logger"
	"time"
)

var RequestCount int
var ServiceErrorCount int
var CountFrom time.Time

func RequestTracer(r *gin.Context) {
	uu := uuid.New()
	r.Set("request_id", uu.String())
	if time.Duration(time.Now().Sub(CountFrom).Hours()) > 1 {
		logger.InfoLog(fmt.Sprintf("Request count: %d \n 500 Service Error count: %d \n Counting From %v \n",
			RequestCount, ServiceErrorCount, CountFrom.String()))
		RequestCount = 0
		ServiceErrorCount = 0
		CountFrom = time.Now()
	}
	RequestCount++
	r.Next()
}
