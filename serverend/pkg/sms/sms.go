package sms

import (
	"context"
	"errors"
)

var (
	ErrRateLimited = errors.New("请一分钟后再操作发送验证码")
)

type Sms interface {
	Send(ctx context.Context, tplId string, args map[string]string, phones []string) error
}
