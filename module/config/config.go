package config

import (
	"fmt"
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
	//读取开关配置
	//开关有2个作用，1是否读取配置文件，2 是否初始化相应的类库
	var switchConfig Switch
	fileContent, err := ioutil.ReadFile(filepath.Join(path, "switch.yaml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileContent, &switchConfig)
	if err != nil {
		panic(err)
	}

	if switchConfig.Mysql {
		var mysqlConfigs []MysqlConfig
		fileContent, err = ioutil.ReadFile(filepath.Join(path, "mysql.yaml"))
		if err != nil {
			panic(err)
		}
		fmt.Println(fileContent)
		err = yaml.Unmarshal(fileContent, &mysqlConfigs)
		//for _, mysqlConfig := range mysqlConfigs {
		//	fmt.Printf("ip is %v\n", mysqlConfig.Ip)
		//}
	}
	return &config
}
