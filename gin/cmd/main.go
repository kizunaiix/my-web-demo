package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	api "ki9.com/gin_demo/routers"
)

// 😀🆒🐉👌
func main() {
	r := gin.Default()

	//CORS策略
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:80", "http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	// r.LoadHTMLGlob("../resources/templates/*.html")
	// r.Static("static", "../resources/static")

	//注册路由
	api.RegisterRouters(r)

	r.Run(":9000")

}
