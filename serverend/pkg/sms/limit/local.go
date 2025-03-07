package limit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"golangdemo/component/limiter"
	redislimiter "golangdemo/component/limiter/redis"
	"microservicedemo/pkg/sms"
	"microservicedemo/pkg/sms/local"
	"time"
)

type SmsLocalLimitService struct {
	service sms.Sms
	limiter limiter.Limiter
}

func NewSmsLocalLimitService(r redis.Cmdable, w time.Duration, num int) *SmsLocalLimitService {
	l := redislimiter.NewSlideWindowLimiter(r, w, num)
	return &SmsLocalLimitService{
		service: &local.SmsLocalService{},
		limiter: l,
	}
}

func (t *SmsLocalLimitService) Send(ctx context.Context, tplId string, args map[string]string, phones []string) error {
	key := "sms-local-limit:" + tplId + ":" + phones[0]
	isLimited, err := t.limiter.IsLimited(ctx, key)
	if err != nil {
		return err
	}
	if isLimited {
		return sms.ErrRateLimited
	}
	return t.service.Send(ctx, tplId, args, phones)
}
