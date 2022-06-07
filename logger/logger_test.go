package logger

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		logLevel          string
		development       bool
		disableCaller     bool
		disableStacktrace bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"zap logger",
			args{
				logLevel:          "debug",
				development:       true,
				disableCaller:     true,
				disableStacktrace: true,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLogger(tt.args.logLevel, tt.args.development, tt.args.disableCaller, tt.args.disableStacktrace)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.Logger.Debug("hello debug")
		})
	}
}
