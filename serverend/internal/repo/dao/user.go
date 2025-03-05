package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"microservicedemo/internal/repo/dao/model"
	"time"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
	ErrUserDuplicate  = errors.New("用户名、手机号、邮箱，重复")
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (t *UserDao) Insert(ctx context.Context, u model.User) error {
	now := time.Now().UnixMilli()
	u.CreatedAt = now
	u.UpdatedAt = now
	err := t.db.WithContext(ctx).Create(&u).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			const uniqueConflictsErrno uint16 = 1062
			if mysqlErr.Number == uniqueConflictsErrno {
				//唯一索引冲突
				return ErrUserDuplicate
			}
		}
		return err
	}
	return nil
}

func (t *UserDao) FindByUsername(ctx context.Context, name string) (model.User, error) {
	var u model.User
	err := t.db.WithContext(ctx).Where("username = ?", name).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, ErrRecordNotFound
	}
	return u, err
}

func (t *UserDao) FindByUserId(ctx context.Context, id int64) (model.User, error) {
	var u model.User
	err := t.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, ErrRecordNotFound
	}
	return u, err
}
