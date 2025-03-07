//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"microservicedemo/internal/repo"
	"microservicedemo/internal/repo/dao"
	"microservicedemo/internal/service"
	"microservicedemo/internal/service/verifycode"
	"microservicedemo/internal/web"
	"microservicedemo/ioc"
	emaillocal "microservicedemo/pkg/email/local"
)

func InitServer() *gin.Engine {
	wire.Build(
		ioc.InitGorm,
		ioc.InitRedis,

		//dao部分
		dao.NewUserDao,
		dao.NewVerifyCodeDao,

		//cache部分

		//repo部分
		repo.NewUserRepo,
		repo.NewVerifyCodeRepo,

		//service部分
		service.NewUserService,

		emaillocal.NewEmailLocalService,
		verifycode.NewEmailVerifyCodeSender,
		verifycode.NewEmailVerifyCodeService,

		ioc.InitSmsService,
		verifycode.NewSmsVerifyCodeSender,
		verifycode.NewSmsVerifyCodeService,

		//web部分
		web.NewUserHandler,

		//最终的服务
		ioc.InitWebServer,
	)
	return &gin.Engine{}
}
