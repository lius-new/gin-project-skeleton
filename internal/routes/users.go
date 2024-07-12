package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lius-new/ln-blog/internal/app"
	"github.com/lius-new/ln-blog/internal/lib"
)

var userRoute = app.Aapplication.CreateGroup("/users")

func init() {
	userRoute.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, lib.ResponseClient.Success("查询指定用户", struct{}{}))
	})
}
