package logger

import (
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	*zap.Logger
}

var levels = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func NewLogger(
	logLevel string,
	development bool,
	disableCaller bool,
	disableStacktrace bool,
) (*Log, error) {
	var encoding string
	if development {
		encoding = "console"
	} else {
		encoding = "json"
	}

	level := zapcore.DebugLevel
	for k, v := range levels {
		if k == strings.ToLower(logLevel) {
			level = v
		}
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       development,
		DisableCaller:     disableCaller,
		DisableStacktrace: disableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
			Hook: func(zapcore.Entry, zapcore.SamplingDecision) {
			},
		},
		Encoding:         encoding,
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{},
	}
	// overwrite encoder config
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// build logger
	logger, err := config.Build()
	if err != nil {
		return nil, errors.Wrap(err, "error building zap logger")
	}
	return &Log{
		logger,
	}, nil
}
