package config

type rdsConfig interface {
	GetIP() string
	GetPort() string
	GetPass() string
	GetMaxOpen() int
}

type defaultRedisConfig struct {
	IP      string `yml:"ip"`
	Port    string `yml:"port"`
	Pass    string `yml:"pass"`
	MaxOpen int    `yml:"max_open"`
}

func (m defaultRedisConfig) GetIP() string {
	return m.IP
}
func (m defaultRedisConfig) GetPort() string {
	return m.Port
}
func (m defaultRedisConfig) GetPass() string {
	return m.Pass
}
func (m defaultRedisConfig) GetMaxOpen() int {
	return m.MaxOpen
}
