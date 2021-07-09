package hutils

import (
	"hutils/boot"
	moduleConfig "hutils/module/config"
	"hutils/module/db"
	utilsLog "hutils/module/log"
	"sync"
)

var (
	m sync.Mutex
)

//Init 初始化
func Init(path string) {

	m.Lock()
	defer m.Unlock()
	//初始化配置
	config := moduleConfig.Init()
	config.Init(path)
	//配置文件
	boot.Config = config
	//初始化日志
	boot.Log = utilsLog.InitLogger(config.GetLogConfig().ZapLog)
	//初始化数据库
	boot.Mysql = db.InitMysql(config.GetMysqlConfig())
}
