package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"microservicedemo/internal/domain"
	"microservicedemo/internal/repo/dao"
	"microservicedemo/internal/repo/dao/model"
)

var (
	ErrUserDuplicate = dao.ErrUserDuplicate
	ErrUserNotFound  = dao.ErrRecordNotFound
)

type UserRepo struct {
	dao *dao.UserDao
}

func NewUserRepo(dao *dao.UserDao) *UserRepo {
	return &UserRepo{dao: dao}
}

func (t *UserRepo) UserDomainToModel(d domain.User) model.User {
	return model.User{
		Username: d.Username,
		Password: d.Password,
		Email: sql.NullString{
			String: d.Email,
			Valid:  d.Email != "",
		},
		Phone: sql.NullString{
			String: d.Phone,
			Valid:  d.Phone != "",
		},
	}
}

func (t *UserRepo) UserModelToDomain(m model.User) domain.User {
	return domain.User{
		Id:       m.Id,
		Username: m.Username,
		Password: m.Password,
	}
}

func (t *UserRepo) Create(ctx context.Context, u domain.User) error {
	_, err := t.dao.Insert(ctx, t.UserDomainToModel(u))
	if errors.Is(err, dao.ErrUserDuplicate) {
		return ErrUserDuplicate
	}
	return err
}

func (t *UserRepo) FindByUserId(ctx context.Context, id int64) (domain.User, error) {
	m, err := t.dao.FindByUserId(ctx, id)
	if errors.Is(err, dao.ErrRecordNotFound) {
		return domain.User{}, ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	return t.UserModelToDomain(m), nil
}

func (t *UserRepo) FindByUsername(ctx context.Context, name string) (domain.User, error) {
	m, err := t.dao.FindByUsername(ctx, name)
	if errors.Is(err, dao.ErrRecordNotFound) {
		return domain.User{}, ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	return t.UserModelToDomain(m), nil
}

func (t *UserRepo) FindOrCreateByEmail(ctx context.Context, email string) (domain.User, error) {
	m, err := t.dao.FindByEmail(ctx, email)
	if err == nil {
		//有数据
		return t.UserModelToDomain(m), nil
	}
	if !errors.Is(err, dao.ErrRecordNotFound) {
		//系统错误
		return domain.User{}, err
	}
	//查空
	d := domain.User{
		Username: "用户" + email,
		Email:    email,
	}
	m, err = t.dao.Insert(ctx, t.UserDomainToModel(d))
	if err == nil {
		return t.UserModelToDomain(m), nil
	}
	if errors.Is(err, dao.ErrUserDuplicate) {
		return domain.User{}, ErrUserDuplicate
	}
	return domain.User{}, err
}

func (t *UserRepo) FindOrCreateByPhone(ctx *gin.Context, phone string) (domain.User, error) {
	m, err := t.dao.FindByPhone(ctx, phone)
	if err == nil {
		//有数据
		return t.UserModelToDomain(m), nil
	}
	if !errors.Is(err, dao.ErrRecordNotFound) {
		//系统错误
		return domain.User{}, err
	}
	//查空
	d := domain.User{
		Username: "用户" + phone,
		Phone:    phone,
	}
	m, err = t.dao.Insert(ctx, t.UserDomainToModel(d))
	if err == nil {
		return t.UserModelToDomain(m), nil
	}
	if errors.Is(err, dao.ErrUserDuplicate) {
		return domain.User{}, ErrUserDuplicate
	}
	return domain.User{}, err
}
