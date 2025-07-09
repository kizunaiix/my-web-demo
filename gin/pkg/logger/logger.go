package logger //TODO: 让GPT再review一下

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init() {
	var err error
	Logger, err = NewLogger(os.Getenv("ENV"))
	if err != nil {
		log.Fatal("Failed to initialize logger:", zap.Error(err))
	}
}

func NewLogger(env string) (ZapLogger *zap.Logger, err error) {

	switch env {

	case "dev":

		ZapLogger, err = zap.NewDevelopment()

	case "prod":

		ZapLogger, err = zap.NewProduction()

	default:

		logCfg := zap.NewProductionConfig()
		logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		ZapLogger, err = logCfg.Build()

	}

	if err != nil {
		ZapLogger.Error("Logger initializing failed", zap.String("env", env))
	} else {
		ZapLogger.Info("Logger initialized", zap.String("env", env))
	}

	return

}

func LoggerMiddleWare(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		//开始计时
		start := time.Now()

		c.Next() // 调用下一个处理器

		//结束计时
		end := time.Now()
		duration := end.Sub(start)

		l.Info("Request Completed",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("duration", duration),
			zap.Int("status", c.Writer.Status()),
			zap.Int("size", c.Writer.Size()),
		)

	}
}
