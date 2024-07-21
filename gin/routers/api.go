package api

import (
	"github.com/gin-gonic/gin"
	"ki9.com/gin_demo/controllers"
)

func RegisterRouters(r *gin.Engine) {

	//TOswag：

	r.GET("/api", controllers.Welcome)

	//TOswag：
	r.POST("/api/login", controllers.Login)
}
