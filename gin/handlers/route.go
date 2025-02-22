package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &struct{ SwaggerDocs string }{SwaggerDocs: "http://localhost:9000/"})
	})

	r.GET("/api/welcome", Welcome)

	r.POST("/api/posttest", PostTest)

}
