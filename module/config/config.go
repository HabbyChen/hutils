package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	logConfig   LogConfig
	mysqlConfig []MysqlConfig
	appConfig   ServiceConfig
}

func Init() *Config {
	return &Config{}
}

func (config Config) GetLogConfig() LogConfig {
	return config.logConfig
}
func (config Config) GetMysqlConfig() []MysqlConfig {
	return config.mysqlConfig
}
func (config Config) GetAppConfig() ServiceConfig {
	return config.appConfig
}

//Need 初始化实际需要的
func (config *Config) Init(path string) {

	//初始化app基础配置
	fileContent, err := ioutil.ReadFile(filepath.Join(path, "service.yml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileContent, &config.appConfig)
	if err != nil {
		panic(err)
	}
	//读取log配置
	fileContent, err = ioutil.ReadFile(filepath.Join(path, "log.yaml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileContent, &config.logConfig)
	if err != nil {
		panic(err)
	}

	//读取mysql配置
	fileContent, err = ioutil.ReadFile(filepath.Join(path, "mysql.yaml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileContent, &config.mysqlConfig)
	if err != nil {
		panic(err)
	}

}

//func (config Config) CusInit(path string) map[string]interface{} {
//	fileContent, err := ioutil.ReadFile("yml/private.yml")
//	var aaa map[string]interface{}
//	err = yaml.Unmarshal(fileContent, &aaa)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("获取到的配置是 %+v", aaa)
//	return aaa
//}
