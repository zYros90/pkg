package logger

import (
	"context"
	"fmt"
	"time"

	gormLogger "gorm.io/gorm/logger"
	gormUtils "gorm.io/gorm/utils"
)

const (
	logTitle      = "[gorm] "
	sqlFormat     = logTitle + "%s"
	messageFormat = logTitle + "%s, %s"
	errorFormat   = logTitle + "%s, %s, %s"
	slowThreshold = 200
)

// LogMode: The log level of gorm logger is overwrited by the log level of Zap logger.
func (l *Log) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

// Info prints a information log.
func (l *Log) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Sugar().Infof(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Warn prints a warning log.
func (l *Log) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Sugar().Warnf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Error prints a error log.
func (l *Log) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Sugar().Errorf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Trace prints a trace log such as sql, source file and error.
func (l *Log) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	switch {
	case err != nil:
		sql, _ := fc()
		l.Sugar().Errorf(errorFormat, gormUtils.FileWithLineNum(), err, sql)
	case elapsed > slowThreshold*time.Millisecond && slowThreshold*time.Millisecond != 0:
		sql, _ := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		l.Sugar().Warnf(errorFormat, gormUtils.FileWithLineNum(), slowLog, sql)
	default:
		sql, _ := fc()
		l.Sugar().Debugf(sqlFormat, sql)
	}
}
