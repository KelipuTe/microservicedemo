package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type LoginCheckMidBuilder struct {
	ignorePath map[string]string
}

func NewLoginCheckMidBuilder() *LoginCheckMidBuilder {
	return &LoginCheckMidBuilder{
		ignorePath: make(map[string]string),
	}
}

func (t *LoginCheckMidBuilder) AddIgnorePath(path []string) *LoginCheckMidBuilder {
	for _, val := range path {
		t.ignorePath[val] = val
	}
	return t
}

func (t *LoginCheckMidBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//不需要登录的路由
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
		//检查一下userId对不对，这样业务代码就可以不用判断了
		if _, ok := userId.(int64); !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//每1个小时刷新一次
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
