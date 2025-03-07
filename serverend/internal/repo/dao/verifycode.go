package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"microservicedemo/internal/repo/dao/model"
	"time"
)

type VerifyCodeDao struct {
	db *gorm.DB
}

func NewVerifyCodeDao(db *gorm.DB) *VerifyCodeDao {
	return &VerifyCodeDao{db: db}
}

func (t *VerifyCodeDao) Insert(ctx context.Context, m model.VerifyCode) error {
	now := time.Now().UnixMilli()
	m.CreatedAt = now
	m.UpdatedAt = now
	return t.db.WithContext(ctx).Create(&m).Error
}

func (t *VerifyCodeDao) Find(ctx context.Context, biz, target string) (model.VerifyCode, error) {
	m := model.VerifyCode{}
	err := t.db.WithContext(ctx).
		Where("biz=? and target=?", biz, target).
		Order("created_at desc").First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, ErrRecordNotFound
	}
	return m, err
}
