package main

import (
	"github.com/image-transform/internal/myconf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(conf *myconf.MyConf) (*zap.Logger, error) {
	z, err := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.DebugLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
		OutputPaths: []string{"stdout"},
	}.Build()
	if err != nil {
		return nil, err
	}
	return z, nil
}
