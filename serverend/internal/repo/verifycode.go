package repo

import (
	"context"
	"errors"
	"microservicedemo/internal/domain"
	"microservicedemo/internal/repo/dao"
	"microservicedemo/internal/repo/dao/model"
	"time"
)

var ErrVerifyCodeNotFound = dao.ErrRecordNotFound

type VerifyCode struct {
	dao *dao.VerifyCodeDao
}

func NewVerifyCodeRepo(dao *dao.VerifyCodeDao) *VerifyCode {
	return &VerifyCode{dao: dao}
}

func (t *VerifyCode) Save(ctx context.Context, biz, target, code string) error {
	expiresAt := time.Now().Add(5 * time.Minute).UnixMilli()
	m := model.VerifyCode{
		Biz:       biz,
		Target:    target,
		Code:      code,
		ExpiresAt: expiresAt,
	}
	return t.dao.Insert(ctx, m)
}

func (t *VerifyCode) Find(ctx context.Context, biz, target string) (domain.VerifyCode, error) {
	m, err := t.dao.Find(ctx, biz, target)
	if errors.Is(err, dao.ErrRecordNotFound) {
		return domain.VerifyCode{}, ErrVerifyCodeNotFound
	}
	if err != nil {
		return domain.VerifyCode{}, err
	}
	return t.VerifyCodeModelToDomain(m), nil
}

func (t *VerifyCode) VerifyCodeModelToDomain(m model.VerifyCode) domain.VerifyCode {
	return domain.VerifyCode{
		Id:        m.Id,
		Biz:       m.Biz,
		Target:    m.Target,
		Code:      m.Code,
		ExpiresAt: m.ExpiresAt,
	}
}
