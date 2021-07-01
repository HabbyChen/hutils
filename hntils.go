package hutils

import (
	moduleConfig "hutils/module/config"

	"go.uber.org/zap"
)

var (
	//Config 配置
	config *moduleConfig.Config
	//日志
	zapLog *zap.SugaredLogger
)

//Init 初始化
func Init(path string) {
	//初始化配置
	config = moduleConfig.Init()
	config.Need(path)
	//初始化日志

	//zapLog = log.InitLogger()
}

//Config 获取配置
func Config() *moduleConfig.Config {
	return config
}

//Log 获取日志
func Log() *zap.SugaredLogger {
	return zapLog
}
