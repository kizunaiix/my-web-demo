package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	_ "ki9.com/gin_demo/cmd/docs" // swagger:这里要用你的实际 `docs` 路径
	"ki9.com/gin_demo/internal/conf"
	"ki9.com/gin_demo/internal/router"
	"ki9.com/gin_demo/pkg/logger"
)

var cfg = conf.Conf{}

// @title 习惯养成 API 文档
// @version 0.3.2
// @description 这是一个基于 Gin 和 RPC 风格的习惯养成 API。应当只有开发环境下能看到该文档。
// @host localhost:9000
// @BasePath /api
func main() {

	// 初始化 logger
	err := logger.Init()
	if err != nil {
		log.Fatal("Logger initializing failed", zap.Error(err))
	}

	defer logger.Logger.Sync() // 确保日志在程序结束时被写入

	// 加载配置
	//加载conf文件的内容 TODO: 封装成func loadConfig(path string) (*conf.Conf, error)
	confFile, err := os.ReadFile("../conf/ginconf-prod.yml")
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(confFile, &cfg); err != nil {
		panic(err)
	}

	// 4. 初始化数据库

	// 5. 初始化 Gin
	r := gin.New()
	r.Use(
		logger.LoggerMiddleware(logger.Logger),
		// gin.Logger(),
		gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:80", "http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	// 6. 注册路由
	router.RegisterRouters(r)

	// 7. 注册 Swagger
	// 如果ENV这个环境变量里不包括prod的话就加上swagger doc
	if !strings.Contains(os.Getenv("ENV"), "prod") {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 8. 启动服务
	logger.Logger.Info("Gin server starting...",
		zap.String("port", "9000"),
		zap.String("env", os.Getenv("ENV")),
	)

	r.Run("0.0.0.0:9000")
}
