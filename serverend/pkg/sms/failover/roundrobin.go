package failover

import (
	"context"
	"errors"
	"fmt"
	"microservicedemo/pkg/sms"
	"sync/atomic"
)

type SmsRoundRobinFailOverService struct {
	services []sms.Sms
	index    uint64
}

func NewSmsRoundRobinFailOverService(sms []sms.Sms) *SmsRoundRobinFailOverService {
	return &SmsRoundRobinFailOverService{
		services: sms,
	}
}

func (t *SmsRoundRobinFailOverService) Send(ctx context.Context, tplId string, args map[string]string, phones []string) error {
	var errStr string
	index := atomic.AddUint64(&t.index, 1)
	length := uint64(len(t.services))
	for i := index; i < index+length; i++ {
		svc := t.services[i%length]
		err := svc.Send(ctx, tplId, args, phones)
		switch {
		case err == nil:
			return nil
		case errors.Is(err, context.Canceled):
		case errors.Is(err, context.DeadlineExceeded):
			return err
		}
		errStr += fmt.Sprintf("第%d次发送失败，原因：%s；", i+1, err.Error())
	}
	return errors.New("本地异常转移短信服务轮询全部失败；" + errStr)
}
