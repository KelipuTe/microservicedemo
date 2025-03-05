package service

import (
	"context"
	"demo-golang/microservice/internal/domain"
	"demo-golang/microservice/internal/repo"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicate         = repo.ErrUserDuplicate
	ErrInvalidUserOrPassword = errors.New("用户名或者密码不对")
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

func (t *UserService) Login(ctx context.Context, u domain.User) (domain.User, error) {
	u2, err := t.repo.FindByUsername(ctx, u.Username)
	if errors.Is(err, repo.ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u2.Password), []byte(u.Password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u2, nil
}
