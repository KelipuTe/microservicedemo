package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golangdemo/component/limiter"
	redislimiter "golangdemo/component/limiter/redis"
	"net/http"
	"time"
)

type RateLimitMidBuilder struct {
	limiter limiter.Limiter
}

func NewRateLimitMidBuilder(r redis.Cmdable, w time.Duration, num int) *RateLimitMidBuilder {
	return &RateLimitMidBuilder{
		limiter: redislimiter.NewSlideWindowLimiter(r, w, num),
	}
}

func (t *RateLimitMidBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		key := "req-limiter:" + ip
		isLimited, err := t.limiter.IsLimited(ctx, key)

		if err != nil {
			//走到这里说明redis报错了，一般是默认限制住
			//但是也有采用放行的方案，尽可能提供服务给用户
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if isLimited {
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		ctx.Next()
	}
}
