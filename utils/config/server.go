package config

type serversConfig interface {
	GetEtcd() string
	GetName() string
	GetAddr() string
}

type defaultServerConfig struct {
	Etcd string `yml:"etcd"`
	Name string `yml:"name"`
	Addr string `yml:"addr"`
}

func (m defaultServerConfig) GetEtcd() string {
	return m.Etcd
}

func (m defaultServerConfig) GetName() string {
	return m.Name
}
func (m defaultServerConfig) GetAddr() string {
	return m.Addr
}
