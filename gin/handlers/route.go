package handlers

import "github.com/gin-gonic/gin"

func RegisterRouters(r *gin.Engine) {

	r.GET("/api", Welcome)

	r.POST("/api/login", Login)

}
