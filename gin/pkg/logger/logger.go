package logger

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(env string) (l *zap.Logger, err error) {

	// env := os.Getenv("ENV")

	switch env {

	case "dev":

		l, err = zap.NewDevelopment()

	case "prod":

		logCfg := zap.NewProductionConfig()
		logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		l, err = logCfg.Build()

	default:
		return nil, errors.New("unknown env, use dev or prod")
	}

	if err != nil {
		l.Error("Logger initializing failed", zap.String("env", env))
	} else {
		l.Info("Logger initializing successed", zap.String("env", env))
	}

	return

}

func LoggerMiddleware(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		//开始计时
		start := time.Now()
		c.Set("logger", l) // 将logger存入context，方便后续调用

		c.Next() // 调用下一个处理器

		//结束计时
		duration := time.Since(start)

		l2, ok := c.Get("logger")
		if !ok {
			l.Error("Logger not found in context")

		}
		l2.(*zap.Logger).Info("Request Completed",
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
