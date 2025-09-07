package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	_ "ki9.com/gin_demo/cmd/docs" // swagger:这里要用你的实际 `docs` 路径
	"ki9.com/gin_demo/internal/conf"
	"ki9.com/gin_demo/internal/server"
	"ki9.com/gin_demo/pkg/logger"
)

var cfg = conf.Conf{}

// @title 习惯养成 API 文档
// @version 0.3.2
// @description 这是一个基于 Gin 和 RPC 风格的习惯养成 API。应当只有开发环境下能看到该文档。
// @host localhost:9000
// @BasePath /api
func main() {

	// --------------------------------------------------------------------
	// 初始化 logger
	err := logger.InitLogger()
	if err != nil {
		log.Fatal("Logger initializing failed", zap.Error(err))
	}

	defer logger.Logger.Sync() // 确保日志在程序结束时被写入

	// --------------------------------------------------------------------
	// 加载配置
	//加载conf文件的内容 TODO: 封装成func loadConfig(path string) (*conf.Conf, error)
	confFile, err := os.ReadFile("../conf/ginconf-prod.yml")
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(confFile, &cfg); err != nil {
		panic(err)
	}

	server.Run()
}
