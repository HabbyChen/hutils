package config

type LogConfig struct {
	ZapLog ZapLogConfig
}

type ZapLogConfig struct {
	LogFileDir string //保存的路径
	CutHours   int    //多久切割一次 为了简便，我们就按照时间切割就好了，就不整那么多杂七杂八的了
	MaxBackups int    //保存的分数
	Loglevel   int8   //错误级别 1-7:debug-panic
	Enable     bool   //是否开启
}
