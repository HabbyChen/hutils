package db

import (
	"fmt"
	moduleConfig "hutils/module/config"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	//注释1：我觉得好神奇的就是postgres必须依赖这个库才可以跑
	//但是官方的文档里面又没有
	_ "github.com/lib/pq"
	//参见注释1

	"gorm.io/gorm"
)

//DbObj db的相关对象
type Obj struct {
	dbConnections map[string]*gorm.DB
}

//DefaultDb 当前项目使用到的数据库名
const (
	DefaultDb = "default"
)

//InitDb 初始化数据库
func InitMysql(configs []moduleConfig.MysqlConfig) *Obj {

	obj := &Obj{}
	obj.dbConnections = make(map[string]*gorm.DB)
	for _, config := range configs {
		connect, err := obj.dbInit(config)
		if err != nil {
			panic(err)
		}
		obj.dbConnections[config.Name] = connect
	}

	return obj
}

//GetDB 注释懒得写，看下面的DB函数
func (d Obj) GetDB(name string) *gorm.DB {
	if len(name) == 0 {
		name = DefaultDb
	}
	conn, ok := d.dbConnections[name]
	if !ok {
		return nil
	}
	return conn
}

//DB 返回一个数据库链接
//这个链接有可能是pgsql的
//也有可能是mysql的或者其他
//本函数是为了简化上面GetDB的传参操作
func (d Obj) DB() *gorm.DB {
	fmt.Println(len(d.dbConnections))
	conn := d.GetDB(DefaultDb)
	//如果没有给默认名字的话，直接返回第一个
	if conn == nil {
		for _, first := range d.dbConnections {
			return first
		}
	}
	return conn
}

func (d *Obj) dbInit(config moduleConfig.MysqlConfig) (*gorm.DB, error) {

	newLogger := &Logger{}
	newLogger.Config.SlowThreshold = 500 * time.Millisecond
	newLogger.Config.LogLevel = logger.Info

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:1qaz2wsx@tcp(172.31.70.111:3306)/cqkydy?charset=utf8&parseTime=True&loc=Local",

		//DefaultStringSize:         256,   // string 类型字段的默认长度
		//DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
	////因为涉及到要使用连接池
	////所以我们的做法是先用sql.Open生成链接
	////然后再gorm.Open灌注到gorm里面去，就是gorm文档中使用已有链接的章节里面的方式
	//var gormDB *gorm.DB
	////sql.Open是支持多种sql数据库的https://github.com/golang/go/wiki/SQLDrivers
	////因为我们主要是mysql和pgsql所以可以这么用
	////但是如果是redis什么的也放里面就不知道咋整了，可能会单独做吧
	//db, err := sql.Open(database.Dialect, database.Dsn)
	//log.Printf("\n开始连接数据库:%v,\n%v:%v", key, database.Dialect, database.Dsn)
	//if err != nil {
	//	//只要一个数据库连接失败，则启动失败
	//	log.Printf("错误信息：%v", err)
	//	panic("错误:" + database.Dialect + ":" + database.Dsn)
	//}
	//switch database.Dialect {
	//case "mysql":

	//case "postgres":
	//	gormDB, err = gorm.Open(postgres.New(postgres.Config{
	//		Conn: db,
	//	}), &gorm.Config{Logger: newLogger})
	//default:
	//	panic("不能识别的数据库类别")
	//}
	//if err != nil {
	//	panic("gorm初始化失败" + database.Dialect + "：" + database.Dsn + err.Error())
	//}
	////设置连接池
	//db.SetMaxIdleConns(database.MaxOpenConn)
	//db.SetMaxOpenConns(database.MaxIdleConn)
	//db.SetConnMaxLifetime(time.Duration(database.ConnMaxLifetime) * time.Second)
	////装入一个连接数组，用的时候从里面去对应的即可
	//dbConnections[key] = gormDB
	//d.dbConnections = dbConnections
	//log.Printf("MaxOpenConn:%v\n", database.MaxOpenConn)
	//log.Printf("MaxIdleConn:%v\n", database.MaxIdleConn)
	//log.Printf("数据库：%v，初始化完毕\n", key)
}
