package repo

import (
	"context"
	"demo-golang/microservice/internal/domain"
	"demo-golang/microservice/internal/repo/dao"
	"demo-golang/microservice/internal/repo/dao/model"
	"errors"
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

func (t *UserRepo) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	userModel, err := t.dao.FindByUsername(ctx, username)
	if errors.Is(err, dao.ErrRecordNotFound) {
		return domain.User{}, ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	return t.UserModelToDomain(userModel), nil
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
