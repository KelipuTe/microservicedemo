package tencent

import (
	"context"
	tencent "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"log"
)

type SmsTencentService struct {
	client    *tencent.Client
	appId     string
	appSecret string
}

func (t *SmsTencentService) Send(ctx context.Context, tplId string, args map[string]string, phones []string) error {
	log.Println("调用腾讯云的短信SDK", tplId, args, phones)
	return nil
}
