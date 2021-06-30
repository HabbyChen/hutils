package config

type ZapLogConfig struct {
	Development string `yml:"development"`
	LogFileDir  string `yml:"logFileDir"`
	AppName     string `yml:"appName"`
	MaxSize     int    `yml:"maxSize"`
	MaxBackups  int    `yml:"maxBackups"`
	MaxAge      int    `yml:"maxAge"`
	Dc          int64  `yml:"dc"`
}
