package config

type MysqlConfig struct {
	Name string `yml:"name"`
	Ip   string `yml:"ip"`
	//Port     int    `yml:"port"`
	//User     string `yml:"user"`
	//Pwd      string `yml:"pwd"`
	//DbName   string `yml:"dbName"`
	//PoolSize int    `yml:"poolSize"`
}
