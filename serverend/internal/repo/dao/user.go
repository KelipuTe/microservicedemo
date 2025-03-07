package dao

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"microservicedemo/internal/repo/dao/model"
	"time"
)

var (
	ErrUserDuplicate = errors.New("用户名、手机号、邮箱，重复")
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (t *UserDao) Insert(ctx context.Context, m model.User) (model.User, error) {
	now := time.Now().UnixMilli()
	m.CreatedAt = now
	m.UpdatedAt = now
	err := t.db.WithContext(ctx).Create(&m).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			const uniqueConflictsErrno uint16 = 1062
			if mysqlErr.Number == uniqueConflictsErrno {
				//唯一索引冲突
				return model.User{}, ErrUserDuplicate
			}
		}
		return model.User{}, err
	}
	return m, nil
}

func (t *UserDao) FindByUserId(ctx context.Context, id int64) (model.User, error) {
	m := model.User{}
	err := t.db.WithContext(ctx).Where("id=?", id).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, ErrRecordNotFound
	}
	return m, err
}

func (t *UserDao) FindByUsername(ctx context.Context, name string) (model.User, error) {
	m := model.User{}
	err := t.db.WithContext(ctx).Where("username=?", name).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, ErrRecordNotFound
	}
	return m, err
}

func (t *UserDao) FindByEmail(ctx context.Context, email string) (model.User, error) {
	m := model.User{}
	err := t.db.WithContext(ctx).Where("email=?", email).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, ErrRecordNotFound
	}
	return m, err
}

func (t *UserDao) FindByPhone(ctx *gin.Context, phone string) (model.User, error) {
	m := model.User{}
	err := t.db.WithContext(ctx).Where("phone=?", phone).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return m, ErrRecordNotFound
	}
	return m, err
}
