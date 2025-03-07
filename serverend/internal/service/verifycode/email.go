package verifycode

import (
	"context"
	"microservicedemo/pkg/email"
)

type EmailVerifyCodeSender struct {
	email email.Email
}

func NewEmailVerifyCodeSender(e email.Email) *EmailVerifyCodeSender {
	return &EmailVerifyCodeSender{
		email: e,
	}
}

func (t *EmailVerifyCodeSender) Send(ctx context.Context, target, code string) error {
	return t.email.Send(ctx, map[string]string{"code": code}, []string{target})
}
