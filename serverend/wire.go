//go:build wireinject

package main

import (
	"demo-golang/microservice/internal/repo"
	"demo-golang/microservice/internal/repo/dao"
	"demo-golang/microservice/internal/service"
	"demo-golang/microservice/internal/web"
	"demo-golang/microservice/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
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
