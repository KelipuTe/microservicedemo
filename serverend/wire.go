//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"microservicedemo/internal/repo"
	"microservicedemo/internal/repo/dao"
	"microservicedemo/internal/service"
	"microservicedemo/internal/web"
	"microservicedemo/ioc"
)

func InitServer() *gin.Engine {
	wire.Build(
		ioc.InitGorm,

		//ioc.InitRedis,

		dao.NewUserDao,

		repo.NewUserRepo,

		service.NewUserService,

		web.NewUserHandler,

		ioc.InitWebServer,
	)
	return &gin.Engine{}
}
