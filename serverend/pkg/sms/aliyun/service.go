package aliyun

import (
	"context"
	aliyun "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"log"
)

type SmsAliyunService struct {
	client    *aliyun.Client
	appId     string
	appSecret string
}

func (t *SmsAliyunService) Send(ctx context.Context, tplId string, args map[string]string, phones []string) error {
	log.Println("调用阿里云的短信SDK", tplId, args, phones)
	return nil
}
