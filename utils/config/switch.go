package config

// Switch 所有配置开关 配置
type Switch struct {
	Log   bool `yaml:"log"`
	Mysql bool `yaml:"mysql"`
	Nacos bool `yaml:"nacos"`
}
