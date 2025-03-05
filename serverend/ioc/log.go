package ioc

import (
	"go.uber.org/zap"
	"microservicedemo/pkg/logger"
)

func InitLog() logger.Logger {
	cfg := zap.NewDevelopmentConfig()
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
