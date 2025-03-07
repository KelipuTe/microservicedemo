package ioc

import (
	"github.com/redis/go-redis/v9"
	"microservicedemo/pkg/sms"
	smslocal "microservicedemo/pkg/sms/limit"
	"time"
)

func InitSmsService(r redis.Cmdable) sms.Sms {
	return smslocal.NewSmsLocalLimitService(r, time.Minute, 1)
}
