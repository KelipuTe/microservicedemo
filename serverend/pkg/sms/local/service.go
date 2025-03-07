package local

import (
	"context"
	"log"
	"microservicedemo/pkg/sms"
)

type SmsLocalService struct {
}

func NewLocalService() sms.Sms {
	return &SmsLocalService{}
}

func (t *SmsLocalService) Send(ctx context.Context, tplId string, args map[string]string, phones []string) error {
	log.Println("调用本地的短信SDK", tplId, args, phones)
	return nil
}
