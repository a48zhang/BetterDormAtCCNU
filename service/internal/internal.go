package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"main/dao"
)

func WriteError(ctx context.Context, req string, data gin.H) {
	info := dao.LogInfo{
		RequestID: req,
		Data:      data,
	}
	info.Insert(ctx)
}
