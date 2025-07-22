package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ki9.com/gin_demo/internal/helloworld/greet"
	"ki9.com/gin_demo/internal/helloworld/welcome"
	"ki9.com/gin_demo/internal/task"
)

func RegisterRouters(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &struct{ SwaggerDocs string }{SwaggerDocs: "http://localhost:9000/"})
	})

	r.GET("/api/welcome", welcome.Welcome)

	r.POST("/api/posttest", greet.Greet)

	r.POST("/api/handle-task", task.HandleTask)
}
