package log

import (
	"hutils/module/config"
	"io"
	"time"

	rotates "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//InitLogger 初始化日志
func InitLogger(logConfig config.ZapLogConfig) *zap.SugaredLogger {
	encoder := getEncoder()

	//两个interface,判断日志等级
	//info以下归到info日志
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})
	//warn及以上归到warn日志
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoWriter := getLogWriter("./logs/info")
	warnWriter := getLogWriter("./logs/warn")

	//创建zap.Core,for logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoWriter, infoLevel),
		zapcore.NewCore(encoder, warnWriter, warnLevel),
	)
	//生成Logger
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

//getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//得到LogWriter
func getLogWriter(filePath string) zapcore.WriteSyncer {
	warnIoWriter := getWriter(filePath)
	return zapcore.AddSync(warnIoWriter)
}

//日志文件切割
func getWriter(filename string) io.Writer {
	// 保存日志30天，每24小时分割一次日志

	//保存日志30天，每1分钟分割一次日志
	hook, err := rotates.New(
		filename+"_%Y%m%d.log",
		//坐软链
		rotates.WithLinkName(filename),
		//保存多少条记录
		rotates.WithRotationCount(15),
		//多少时间切割一次
		rotates.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
