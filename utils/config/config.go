package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var m sync.Mutex

type Config struct {
	MysqlConfig []mysqlConfig
	RedisConfig []redisConfig
}

//InitNeed 初始化实际需要的
func InitNeed(path string) {
	m.Lock()
	defer m.Unlock()
	//读取开关配置
	//开关有2个作用，1是否读取配置文件，2 是否初始化相应的类库
	var switchConfig Switch
	fmt.Println(filepath.Join(path, "switch.yaml"))
	fileContent, err := ioutil.ReadFile(filepath.Join(path, "switch.yaml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(fileContent, &switchConfig)
	if err != nil {
		panic(err)
	}

	if switchConfig.Mysql == true {
		var mysqlConfigs []mysqlConfig
		fileContent, err = ioutil.ReadFile(filepath.Join(path, "mysql.yaml"))
		if err != nil {
			panic(err)
		}
		fmt.Println(fileContent)
		err = yaml.Unmarshal(fileContent, &mysqlConfigs)
		for _, mysqlConfig := range mysqlConfigs {
			fmt.Printf("ip is %v\n", mysqlConfig.Ip)
		}
	}

	fmt.Print("123123")
}
