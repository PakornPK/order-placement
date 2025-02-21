package logs

import (
	"github.com/PakornPK/order-placement/config"
	"go.uber.org/zap"
)

type Logger interface {
	Info(string, ...zap.Field)
	Error(any, ...zap.Field)
	Warn(string, ...zap.Field)
}

type logger struct {
	logger *zap.Logger
}

func NewLogger(cfg config.AppConfig) (Logger, func() error) {
	var l *zap.Logger
	if cfg.IsLocal() {
		l, _ = zap.NewDevelopment()
	} else {
		l, _ = zap.NewProduction()
	}
	return logger{logger: l}, l.Sync
}

func (l logger) Info(msg string, fields ...zap.Field) {
	l.logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func (l logger) Error(msg any, fields ...zap.Field) {
	switch m := msg.(type) {
	case string:
		l.logger.WithOptions(zap.AddCallerSkip(1)).Error(m, fields...)
	case error:
		l.logger.WithOptions(zap.AddCallerSkip(1)).Error(m.Error(), fields...)
	}
}

func (l logger) Warn(msg string, fields ...zap.Field) {
	l.logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}
