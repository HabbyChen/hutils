package boot

import (
	"hutils/boot/interfaces"
)

//Mysql 数据库相关操作的实现
var Mysql interfaces.DbInterface

//Pgsql 数据库的相关实现
var Pgsql interfaces.DbInterface

//Log 日志的相关实现
var Log interfaces.LoggerInterface

var Config interfaces.ConfigInterface
