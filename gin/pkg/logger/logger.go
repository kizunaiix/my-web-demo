package logger

import (
	"os"

	"go.uber.org/zap"
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
		defer ZapLogger.Sync()
		ZapLogger.Info("Logger initialized", zap.String("env", os.Getenv("ENV")))

	case "prod":

		ZapLogger, err = zap.NewProduction()
		if err != nil {
			return err
		}
		defer ZapLogger.Sync()
		ZapLogger.Info("Logger initialized", zap.String("env", os.Getenv("ENV")))

	default:
		ZapLogger, err = zap.NewDevelopment()
		if err != nil {
			return err
		}
		defer ZapLogger.Sync()
		ZapLogger.Info("Logger initialized", zap.String("env", os.Getenv("ENV")))
	}

	return nil

}
