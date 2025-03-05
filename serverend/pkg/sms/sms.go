package sms

import "context"

type Sms interface {
	Send(ctx context.Context, tplId string, args map[string]string, phones []string)
}
