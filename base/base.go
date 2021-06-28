package base

import (
	"kitutils/base/config"
	"kitutils/base/db"
	"kitutils/base/tool"
)

//配置文件的目录
func Init(path string) {
	config.Init(path)
	tool.Init()
	db.Init()
}
