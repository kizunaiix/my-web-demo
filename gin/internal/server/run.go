package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"ki9.com/gin_demo/internal/conf"
	"ki9.com/gin_demo/internal/helloworld/greet"
	"ki9.com/gin_demo/internal/helloworld/welcome"
	"ki9.com/gin_demo/internal/middleware/myerr"
	"ki9.com/gin_demo/internal/task"
	"ki9.com/gin_demo/pkg/logger"
)

func Run(l *zap.Logger, cfg *conf.Conf) {
	// --------------------------------------------------------------------
	// 初始化 Gin
	r := gin.New()
	r.Use(
		gin.Recovery(),
		logger.LoggerMiddleware(l),
		myerr.ErrorHandlerMiddleware,
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:80", "http://localhost:3000"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		}),
	)

	// --------------------------------------------------------------------
	// 初始化task模块依赖
	memrepo := task.NewTaskRepositoryMemSlice()
	svc := task.NewTaskService(memrepo)
	th := task.NewTaskHandler(svc)

	// --------------------------------------------------------------------
	// 注册 Swagger
	// 如果ENV这个环境变量里不包括prod的话就加上swagger doc
	if !strings.Contains(os.Getenv("ENV"), "prod") {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// --------------------------------------------------------------------
	// 注册路由
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &struct{ SwaggerDocs string }{SwaggerDocs: "http://localhost:9000/"})
	})
	r.GET("/api/welcome", welcome.Welcome)
	r.POST("/api/posttest", greet.Greet)
	r.POST("/api/handle-task", th.TaskHandlerFunc)

	// --------------------------------------------------------------------
	// 启动服务
	l.Info("Gin server starting...",
		zap.String("port", "9000"),
		zap.String("env", os.Getenv("ENV")),
	)

	r.Run("0.0.0.0:9000")
}
