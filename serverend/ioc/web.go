package ioc

import (
	"demo-golang/microservice/internal/web"
	"demo-golang/microservice/internal/web/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitWebServer(uh *web.UserHandler) *gin.Engine {
	server := gin.Default()

	//store, err := redis.NewStore(16, "tcp", "localhost:6379", "")
	//if err != nil {
	//	panic(err)
	//}
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("ssid", store))

	server.Use(middleware.NewLogMiddlewareBuilder(middleware.DefaultSaveLog).Build())

	loginCheck := middleware.NewLoginCheckMiddlewareBuilder()
	loginCheck.AddIgnorePath([]string{"/user/signup", "/user/login"})
	server.Use(loginCheck.Build())

	uh.RegisterRoutes(server)

	return server
}
