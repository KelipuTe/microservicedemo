package email

import "context"

type Email interface {
	Send(ctx context.Context, args map[string]string, emails []string) error
}
