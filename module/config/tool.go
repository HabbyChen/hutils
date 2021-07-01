package config

type toolLogConfig interface {
	GetDevelopment() bool
	GetLogFileDir() string
	GetAppName() string
	GetMaxSize() int
	GetMaxBackups() int
	GetMaxAge() int
	GetDcId() int64
}

type defaultLogToolConfig struct {
	Development string `yml:"development"`
	LogFileDir  string `yml:"logFileDir"`
	AppName     string `yml:"appName"`
	MaxSize     int    `yml:"maxSize"`
	MaxBackups  int    `yml:"maxBackups"`
	MaxAge      int    `yml:"maxAge"`
	Dc          int64  `yml:"dc"`
}
