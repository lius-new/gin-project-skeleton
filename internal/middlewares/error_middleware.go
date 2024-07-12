package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lius-new/ln-blog/internal/lib"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := recover(); err != nil {
			// TODO: 这里直接返回了错误信息，需要对错误信息再进行处理。
			ctx.JSON(http.StatusInternalServerError, lib.ResponseClient.Error(err))
			return
		}
		ctx.Next()
	}
}
