package ioc

import (
	"github.com/gin-contrib/sessions"
	ssredis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"microservicedemo/internal/web"
	mid "microservicedemo/internal/web/middleware"
	pkgmid "microservicedemo/pkg/gin/middleware"
	"time"
)

func InitWebServer(r redis.Cmdable, uh *web.UserHandler) *gin.Engine {
	server := gin.Default()

	store, err := ssredis.NewStore(
		16, "tcp", "localhost:6379", "",
		[]byte("qqqqqqqqwwwwwwwweeeeeeeerrrrrrrr"),
		[]byte("qqqqqqqqwwwwwwwweeeeeeeerrrrrrrr"),
	)
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("ssid", store))

	server.Use(pkgmid.NewLoggerMidBuilder(pkgmid.DefaultSaveLog).Build())
	server.Use(mid.NewCorsMid())
	server.Use(pkgmid.NewRateLimitMidBuilder(r, 60*time.Second, 200).Build())
	loginCheck := mid.NewLoginCheckMidBuilder()
	loginCheck.AddIgnorePath([]string{
		"/user/signup", "/user/login",
		"/user/signup_email", "/user/login_email",
		"/user/signup_sms", "/user/login_sms",
	})
	server.Use(loginCheck.Build())

	uh.RegisterRoutes(server)

	return server
}
