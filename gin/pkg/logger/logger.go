package logger //TODO: 让GPT再review一下

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init() (err error) {
	Logger, err = NewLogger()
	if err != nil {
		return
	}
	return
}

func NewLogger() (l *zap.Logger, err error) {

	env := os.Getenv("ENV")

	switch env {

	case "dev":

		l, err = zap.NewDevelopment()

	case "prod":

		l, err = zap.NewProduction()

	default:

		logCfg := zap.NewProductionConfig()
		logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		l, err = logCfg.Build()

	}

	if err != nil {
		l.Error("Logger initializing failed", zap.String("env", env))
	} else {
		l.Info("Logger initialized", zap.String("env", env))
	}

	return

}

func LoggerMiddleware(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		//开始计时
		start := time.Now()

		c.Next() // 调用下一个处理器

		//结束计时
		// end := time.Now()
		// duration := end.Sub(start)
		duration := time.Since(start)

		l.Info("Request Completed",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("duration", duration),
			zap.Int("size", c.Writer.Size()),
		)

	}
}
