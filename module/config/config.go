package config

import (
	"io/ioutil"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
)

var m sync.Mutex

type Config struct {
	mysqlConfig []MysqlConfig
	redisConfig []RedisConfig
}

func Init() *Config {
	return &Config{}
}
func (config Config) MysqlConfig() []MysqlConfig {
	return config.mysqlConfig
}

func (config Config) RedisConfig() []RedisConfig {
	return config.redisConfig
}

//Need 初始化实际需要的
func (config Config) Need(path string) *Config {
	m.Lock()
	defer m.Unlock()

	var mysqlConfigs []MysqlConfig
	fileContent, err := ioutil.ReadFile(filepath.Join(path, "mysql.yaml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileContent, &mysqlConfigs)
	if err != nil {
		panic(err)
	}
	//for _, mysqlConfig := range mysqlConfigs {
	//	fmt.Printf("ip is %v\n", mysqlConfig.Ip)
	//}
	return &config
}
