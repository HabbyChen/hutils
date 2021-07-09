package log

import (
	config2 "hutils/boot/interfaces/config"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//InitLogger 初始化日志
func InitLogger(logConfig config2.ZapLogConfig) *zap.SugaredLogger {
	encoder := getEncoder()

	//两个interface,判断日志等级
	//info以下归到info日志
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel /*&& lvl >= logLevel*/
	})
	//warn及以上归到warn日志
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.InfoLevel /*&& lvl >= logLevel*/
	})

	infoWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logConfig.LogFileDir, "info"),
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		LocalTime:  true,
		Compress:   true,
	})
	warnWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logConfig.LogFileDir, "error"),
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		LocalTime:  true,
		Compress:   true,
	})
	//创建zap.Core,for logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoWriter, zapcore.AddSync(os.Stdout)), infoLevel),
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(warnWriter, zapcore.AddSync(os.Stdout)), warnLevel),
	)
	//生成Logger
	return zap.New(core, zap.AddCaller()).Sugar()
}

//getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("[2006-01-02 15:04:05.000]")
	return zapcore.NewConsoleEncoder(encoderConfig)
}
