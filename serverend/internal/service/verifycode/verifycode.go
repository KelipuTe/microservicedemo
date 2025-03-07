package verifycode

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"microservicedemo/internal/repo"
	"time"
)

var (
	ErrVerifyCodeNotFound = errors.New("验证码不存在")
	ErrWrongVerifyCode    = errors.New("验证码错误")
	ErrExpiredVerifyCode  = errors.New("验证码过期")
)

type verifyCodeSender interface {
	Send(ctx context.Context, target, code string) error
}

type verifyCodeService struct {
	repo   *repo.VerifyCode
	sender verifyCodeSender
}

type EmailVerifyCodeService struct {
	*verifyCodeService
}

func NewEmailVerifyCodeService(r *repo.VerifyCode, s *EmailVerifyCodeSender) *EmailVerifyCodeService {
	return &EmailVerifyCodeService{
		verifyCodeService: &verifyCodeService{
			repo:   r,
			sender: s,
		},
	}
}

type SMSVerifyCodeService struct {
	*verifyCodeService
}

func NewSmsVerifyCodeService(r *repo.VerifyCode, s *SmsVerifyCodeSender) *SMSVerifyCodeService {
	return &SMSVerifyCodeService{
		verifyCodeService: &verifyCodeService{
			repo:   r,
			sender: s,
		},
	}
}

func (t *verifyCodeService) Send(ctx context.Context, biz, target string) error {
	code := t.newCode()
	err := t.repo.Save(ctx, biz, target, code)
	if err != nil {
		return err
	}
	return t.sender.Send(ctx, target, code)
}

func (t *verifyCodeService) Verify(ctx context.Context, biz, target, code string) error {
	m, err := t.repo.Find(ctx, biz, target)
	if errors.Is(err, repo.ErrVerifyCodeNotFound) {
		return ErrVerifyCodeNotFound
	}
	if err != nil {
		return err
	}
	if code != m.Code {
		return ErrWrongVerifyCode
	}

	now := time.Now().UnixMilli()
	if now > m.ExpiresAt {
		return ErrExpiredVerifyCode
	}
	return nil
}

func (t *verifyCodeService) newCode() string {
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}
