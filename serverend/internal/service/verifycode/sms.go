package verifycode

import (
	"context"
	"microservicedemo/pkg/sms"
)

type SmsVerifyCodeSender struct {
	sms sms.Sms
}

func NewSmsVerifyCodeSender(sms sms.Sms) *SmsVerifyCodeSender {
	return &SmsVerifyCodeSender{
		sms: sms,
	}
}

func (t *SmsVerifyCodeSender) Send(ctx context.Context, target, code string) error {
	return t.sms.Send(ctx, "0", map[string]string{"code": code}, []string{target})
}
