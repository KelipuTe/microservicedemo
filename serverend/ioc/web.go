package ioc

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"microservicedemo/internal/web"
	mid "microservicedemo/internal/web/middleware"
	pkgmid "microservicedemo/pkg/gin/middleware"
)

func InitWebServer(uh *web.UserHandler) *gin.Engine {
	server := gin.Default()

	store, err := redis.NewStore(
		16, "tcp", "localhost:6379", "",
		[]byte("qqqqqqqqwwwwwwwweeeeeeeerrrrrrrr"),
		[]byte("qqqqqqqqwwwwwwwweeeeeeeerrrrrrrr"),
	)
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("ssid", store))

	server.Use(pkgmid.NewLogMiddlewareBuilder(pkgmid.DefaultSaveLog).Build())
	server.Use(mid.NewCorsMid())
	loginCheck := mid.NewLoginCheckMidBuilder()
	loginCheck.AddIgnorePath([]string{"/user/signup", "/user/login"})
	server.Use(loginCheck.Build())

	uh.RegisterRoutes(server)

	return server
}
