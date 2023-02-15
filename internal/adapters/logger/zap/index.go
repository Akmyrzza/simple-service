package zap

import (
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const callerSkip = 1

type St struct {
	l  *zap.Logger
	sl *zap.SugaredLogger
}

func New(level string, dev bool) *St {
	var cfg zap.Config

	if dev {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
		cfg.Level.SetLevel(getZapLevel(level))
	}

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	l, err := cfg.Build(zap.AddCallerSkip(callerSkip))
	if err != nil {
		log.Fatal(err)
	}

	return &St{
		l:  l,
		sl: l.Sugar(),
	}
}

func getZapLevel(v string) zapcore.Level {
	switch strings.ToLower(v) {
	case "error":
		return zap.ErrorLevel
	case "warn":
		return zap.WarnLevel
	case "info":
		return zap.InfoLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}

func (o *St) Warnw(msg string, args ...any) {
	o.sl.Warnw(msg, args...)
}

func (o *St) Errorw(msg string, err any, args ...any) {
	o.sl.Errorw(msg, err, args)
}

func (o *St) Infow(msg string, args ...any) {
	o.sl.Infow(msg, args...)
}
