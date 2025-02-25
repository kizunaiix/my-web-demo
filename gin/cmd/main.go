package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/yaml.v3"
	_ "ki9.com/gin_demo/cmd/docs" // swagger:这里要用你的实际 `docs` 路径
	"ki9.com/gin_demo/internal/handler"
	"ki9.com/gin_demo/internal/model/conf"
)

var appconf = conf.Conf{}

func init() {

	//加载conf文件的内容
	confFile, err := os.ReadFile("../conf/ginconf-prod.yml")
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(confFile, &appconf); err != nil {
		panic(err)
	}

}

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
	if appconf.SwagEnable {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//注册路由
	handler.RegisterRouters(r)

	r.Run(":9000")

}
