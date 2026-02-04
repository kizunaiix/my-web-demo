package server

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"ki9.com/gin_demo/internal/config"
	"ki9.com/gin_demo/internal/helloworld"
	"ki9.com/gin_demo/internal/middleware/myerr"
	"ki9.com/gin_demo/internal/task"
	"ki9.com/gin_demo/pkg/logging"
)

func Run(l *zap.Logger, cfg config.Conf) {
	//// 初始化 Gin
	r := gin.New()

	r.SetTrustedProxies([]string{"127.0.0.1/32", "192.168.0.0/16"})

	r.Use(
		gin.Recovery(),
		logging.LoggerMiddleware(l),
		myerr.ErrorHandlerMiddleware,
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:80", "http://localhost:3000"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		}),
	)

	//// 初始化task模块依赖
	memrepo := task.NewTaskRepositoryMemSlice()
	svc := task.NewTaskService(memrepo)
	th := task.NewTaskHandler(svc)

	//// 注册 Swagger
	// 如果ENV这个环境变量里不包括prod的话就加上swagger doc
	if !strings.Contains(os.Getenv("ENV"), "prod") {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//// 注册路由
	// 所有不存在的路由返回前端index.html
	r.Static("/my-web-demo", "./public/")

	r.NoRoute(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
			ctx.JSON(404, gin.H{"msg": "api not found"})

		} else {
			ctx.File("../public/index.html")
		}
	})

	api := r.Group("/api/my-web-demo")

	api.GET("/welcome", helloworld.Welcome)
	api.POST("/handle-task", th.TaskHandlerFunc)

	//// 启动服务
	l.Info("Gin server starting...",
		zap.String("port", "9000"),
		zap.String("env", os.Getenv("ENV")),
	)

	r.Run("0.0.0.0:9000")
}
