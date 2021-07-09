package interfaces

import "hutils/boot/interfaces/config"

//ConfigInterface 数据库需要实现的忌口
type ConfigInterface interface {
	GetLogConfig() config.LogConfig
	GetServiceConfig() config.ServiceConfig
}
