package local

import (
	"context"
	"log"
	"microservicedemo/pkg/email"
)

type EmailLocalService struct {
}

func NewEmailLocalService() email.Email {
	return &EmailLocalService{}
}

func (t *EmailLocalService) Send(ctx context.Context, args map[string]string, emails []string) error {
	log.Println("调用本地的邮箱SDK", args, emails)
	return nil
}
