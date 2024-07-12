package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lius-new/ln-blog/internal/app"
	"github.com/lius-new/ln-blog/internal/lib"
)

var articleRoute = app.Aapplication.CreateGroup("/articles")

func init() {
	articleRoute.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, *lib.ResponseClient.Success("查询指定文章", struct{}{}))
	})
}
