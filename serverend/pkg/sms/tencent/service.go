package tencent

import (
	"context"
	tencent "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Service struct {
	client *tencent.Client
}

func (t *Service) Send(ctx context.Context, tplId string, args map[string]string, phones []string) {
	panic("implement me")
}
