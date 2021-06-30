package config

type rdsConfig interface {
	GetIP() string
	GetPort() string
	GetPass() string
	GetMaxOpen() int
}

type RedisConfig struct {
	IP      string `yml:"ip"`
	Port    string `yml:"port"`
	Pass    string `yml:"pass"`
	MaxOpen int    `yml:"max_open"`
}
