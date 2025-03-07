package tencent

import (
	"context"
	"log"
)

type EmailTencentService struct {
	appId     string
	appSecret string
}

func NewEmailTencentService() *EmailTencentService {
	return &EmailTencentService{}
}

func (t *EmailTencentService) Send(ctx context.Context, args map[string]string, emails []string) error {
	log.Println("调用腾讯云的邮箱SDK", args, emails)
	return nil
}
