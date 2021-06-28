package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go.uber.org/zap"
	"kitutils/base/config"
	"kitutils/base/tool"
	"os"
)

//mysql连接池
func initMysql() {
	if len(config.GetMysqlConfig().GetIp())==0{
		tool.NewLogger().Warn("None Mysql Init")
		return
	}
	var err error
	sql := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.GetMysqlConfig().GetUser(), config.GetMysqlConfig().GetPwd(),
		config.GetMysqlConfig().GetIp(), config.GetMysqlConfig().GetPort(), config.GetMysqlConfig().GetDbName())
	tool.NewLogger().Debug("[initMysql] " + sql)
	mysqlEngine, err = xorm.NewEngine("mysql", sql)
	if err != nil {
		tool.NewLogger().Error("[initMysql] "+sql, zap.Error(err))
		os.Exit(0)
	}
	mysqlEngine.SetMaxOpenConns(config.GetMysqlConfig().GetPoolSize())
	mysqlEngine.SetMaxIdleConns(config.GetMysqlConfig().GetPoolSize())
	if err = mysqlEngine.Ping(); err != nil {
		panic(err)
	}
}

func CloseMysqlConnection() {
	_ = mysqlEngine.Close()
}
