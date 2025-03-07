package retry

import (
	"context"
	"errors"
	"fmt"
	"microservicedemo/pkg/sms"
	"microservicedemo/pkg/sms/local"
)

type SmsLocalRetryService struct {
	service  sms.Sms
	tryTimes int
}

func NewSmsLocalRetryService(tryTimes int) *SmsLocalRetryService {
	return &SmsLocalRetryService{
		service: &local.SmsLocalService{},
	}
}

func (t *SmsLocalRetryService) Send(ctx context.Context, tplId string, args map[string]string, phones []string) error {
	var errStr string
	for i := 0; i < t.tryTimes; i++ {
		err := t.service.Send(ctx, tplId, args, phones)
		if err == nil {
			return nil
		}
		errStr += fmt.Sprintf("第%d次发送失败，原因：%s；", i+1, err.Error())
	}
	return errors.New("本地重试短信服务重试全部失败；" + errStr)
}
