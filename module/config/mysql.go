package config

type msqlConfig interface {
	GetIp() string
	GetPort() int
	GetUser() string
	GetPwd() string
	GetDbName() string
	GetPoolSize() int
}

type MysqlConfig struct {
	Ip string `yml:"ip"`
	//Port     int    `yml:"port"`
	//User     string `yml:"user"`
	//Pwd      string `yml:"pwd"`
	//DbName   string `yml:"dbName"`
	//PoolSize int    `yml:"poolSize"`
}
