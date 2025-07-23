package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLogger *zap.Logger
var err error

func InitLogger(env string) error {
	switch env {

	case "dev":

		ZapLogger, err = zap.NewDevelopment()
		if err != nil {
			return err
		}
		ZapLogger.Info("Logger initialized", zap.String("env", os.Getenv("ENV")))

	case "prod":

		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder   //默认是时间戳，手动改成RFC3339格式
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //大写的日志级别
		ZapLogger, err = cfg.Build()
		if err != nil {
			return err
		}
		ZapLogger.Info("Logger initialized", zap.String("env", os.Getenv("ENV")))

	default:
		ZapLogger, err = zap.NewDevelopment()
		if err != nil {
			return err
		}
		ZapLogger.Info("Logger initialized", zap.String("env", os.Getenv("ENV")))
	}

	return nil

}
