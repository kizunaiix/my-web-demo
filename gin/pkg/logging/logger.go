package logging

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(env string) (l *zap.Logger, err error) {

	switch env {

	case "dev":

		l, err = zap.NewDevelopment()

	case "prod":

		logCfg := zap.NewProductionConfig()
		logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		l, err = logCfg.Build()

	default:
		return nil, errors.New("undefined env")
	}

	if err != nil {
		l.Error("Failed to init logger", zap.String("env", env))
	} else {
		l.Info("Logger initialized", zap.String("env", env))
	}

	return

}

func LoggerMiddleware(l *zap.Logger) gin.HandlerFunc { //TODO 以后可以用With()来加上trace_id
	return func(c *gin.Context) {

		//从header拿traceid
		traceID := c.GetHeader("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		LoggerWithTraceID := l.With(zap.String("trace_id", traceID))

		c.Set("logger", LoggerWithTraceID) // 将logger存入context，方便后续调用

		//开始计时
		start := time.Now()

		c.Next()

		//结束计时
		duration := time.Since(start)

		l2 := c.MustGet("logger")

		l2.(*zap.Logger).Info("Request Completed",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("duration", duration),
			zap.Int("body_size", c.Writer.Size()),
			// zap.String("trace_id", traceID),
		)

	}
}
