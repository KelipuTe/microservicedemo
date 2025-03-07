package service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
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

func (t *UserService) Signup(ctx context.Context, d domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	d.Password = string(hash)
	err = t.repo.Create(ctx, d)
	if errors.Is(err, repo.ErrUserDuplicate) {
		return ErrUserDuplicate
	}
	return err
}

func (t *UserService) Login(ctx context.Context, name, pass string) (domain.User, error) {
	d, err := t.repo.FindByUsername(ctx, name)
	if errors.Is(err, repo.ErrUserNotFound) {
		return domain.User{}, ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(pass))
	if err != nil {
		return domain.User{}, ErrWrongPassword
	}
	return d, nil
}

func (t *UserService) LoginEmail(ctx context.Context, email string) (domain.User, error) {
	d, err := t.repo.FindOrCreateByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return d, nil
}

func (t *UserService) FindByUserId(ctx context.Context, id int64) (domain.User, error) {
	d, err := t.repo.FindByUserId(ctx, id)
	if errors.Is(err, repo.ErrUserNotFound) {
		return domain.User{}, ErrUserNotFound
	}
	return d, err
}

func (t *UserService) LoginPhone(ctx *gin.Context, phone string) (domain.User, error) {
	d, err := t.repo.FindOrCreateByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return d, nil
}
