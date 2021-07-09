package config

//MysqlConfig 数据库配置
type MysqlConfig struct {
	Name            string `yml:"name"`     //该链接标志
	IP              string `yml:"ip"`       //ip
	Port            string `yml:"port"`     //ip
	Username        string `yml:"username"` //数据库用户名
	Password        string `yml:"password"` //数据库密码
	Dbname          string `yml:"dbname"`   //dbName
	MaxIdleConn     int    `yml:"max_idle"` //空闲保持的连接数
	MaxOpenConns    int    `yml:"max_open"` //繁忙最大连接数
	ConnMaxLifetime int    `yml:"lifetime"` //长连接保持时间单位分钟
	//Port     int    `yml:"port"`
	//User     string `yml:"user"`
	//Pwd      string `yml:"pwd"`
	//DbName   string `yml:"dbName"`
	//PoolSize int    `yml:"poolSize"`
}
