package db

import (
	"context"
	"hutils/boot"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

type Logger struct {
	Config gormlogger.Config
}

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return &Logger{}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	boot.Log.Debugf(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	boot.Log.Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	boot.Log.Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	//if l.LogLevel <= 0 {
	//	return
	//}
	//elapsed := time.Since(begin)
	//switch {
	//case err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
	//	sql, rows := fc()
	//	l.logger().Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	//case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
	//	sql, rows := fc()
	//	l.logger().Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	//case l.LogLevel >= gormlogger.Info:
	//	sql, rows := fc()
	//	l.logger().Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	//}
}
