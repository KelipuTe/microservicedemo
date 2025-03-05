package ioc

import (
	"demo-golang/microservice/pkg/logger"
	"go.uber.org/zap"
)

func InitLog() logger.Logger {
	cfg := zap.NewDevelopmentConfig()
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
