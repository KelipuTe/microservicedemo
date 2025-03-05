package repo

import (
	"context"
	"errors"
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

func (t *UserRepo) Create(ctx context.Context, u domain.User) error {
	err := t.dao.Insert(ctx, t.UserDomainToModel(u))
	if errors.Is(err, dao.ErrUserDuplicate) {
		return ErrUserDuplicate
	}
	return err
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

func (t *UserRepo) UserDomainToModel(u domain.User) model.User {
	return model.User{
		Username: u.Username,
		Password: u.Password,
	}
}

func (t *UserRepo) UserModelToDomain(u model.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}
