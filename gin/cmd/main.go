package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "ki9.com/gin_demo/docs" // swagger:这里要用你的实际 `docs` 路径
	"ki9.com/gin_demo/handlers"
)

// @title 习惯养成 API 文档
// @version 0.2
// @description 这是一个基于 Gin 和 RPC 风格的习惯养成 API。应当只有开发环境下能看到该文档。
// @host localhost:9000
// @BasePath /api
func main() {
	r := gin.Default()

	//CORS策略
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:80", "http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	// 绑定 Swagger UI路由，仅在go build命令带上-tags dev后生效
	useSwag(r)

	//注册路由
	handlers.RegisterRouters(r)

	r.Run(":9000")

}
