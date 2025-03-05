package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"microservicedemo/internal/domain"
	"microservicedemo/internal/repo"
)

var (
	ErrUserDuplicate = repo.ErrUserDuplicate
	ErrUserNotFound  = errors.New("用户不存在")
	ErrWrongPassword = errors.New("密码错误")
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (t *UserService) Signup(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	err = t.repo.Create(ctx, u)
	if errors.Is(err, repo.ErrUserDuplicate) {
		return ErrUserDuplicate
	}
	return err
}

func (t *UserService) Login(ctx context.Context, name, pass string) (domain.User, error) {
	u, err := t.repo.FindByUsername(ctx, name)
	if errors.Is(err, repo.ErrUserNotFound) {
		return domain.User{}, ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		return domain.User{}, ErrWrongPassword
	}
	return u, nil
}

func (t *UserService) FindByUserId(ctx context.Context, id int64) (domain.User, error) {
	u, err := t.repo.FindByUserId(ctx, id)
	if errors.Is(err, repo.ErrUserNotFound) {
		return domain.User{}, ErrUserNotFound
	}
	return u, err
}
