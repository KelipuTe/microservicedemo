package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type LoginCheckMiddlewareBuilder struct {
	ignorePath map[string]string
}

func NewLoginCheckMiddlewareBuilder() *LoginCheckMiddlewareBuilder {
	return &LoginCheckMiddlewareBuilder{
		ignorePath: make(map[string]string),
	}
}

func (t *LoginCheckMiddlewareBuilder) AddIgnorePath(path []string) *LoginCheckMiddlewareBuilder {
	for _, val := range path {
		t.ignorePath[val] = val
	}
	return t
}

func (t *LoginCheckMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path

		_, ok := t.ignorePath[path]
		if ok {
			return
		}

		//检查登录状态
		sess := sessions.Default(ctx)
		userId := sess.Get("userId")
		if userId == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//每个小时刷新一次
		now := time.Now()
		const updatedAtKey = "updatedAt"
		val := sess.Get(updatedAtKey)
		uat, ok := val.(time.Time)
		if val == nil || !ok || now.Sub(uat) > time.Hour {
			sess.Set("userId", userId)
			sess.Set(updatedAtKey, now)
			sess.Options(sessions.Options{
				MaxAge: 86400,
			})
			err := sess.Save()
			if err != nil {
				log.Println("LoginCheckMiddleware", err)
			}
		}

		ctx.Next()
	}
}
