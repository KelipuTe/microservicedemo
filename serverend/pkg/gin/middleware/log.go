package middleware

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"time"
)

type LogData struct {
	ReqTime    time.Time
	Method     string
	Path       string
	Header     map[string][]string
	ReqData    string
	StatusCode int
	RespData   string
	RespTime   time.Time
	TimeCost   time.Duration
}

type howToSaveLog func(ctx context.Context, l LogData)

func DefaultSaveLog(ctx context.Context, l LogData) {
	log.Println("LogMiddleware", l)
}

type LogMidBuilder struct {
	saveLog  howToSaveLog
	saveReq  bool
	saveResp bool
}

func NewLogMiddlewareBuilder(saveLog howToSaveLog) *LogMidBuilder {
	return &LogMidBuilder{
		saveLog:  saveLog,
		saveReq:  false,
		saveResp: false,
	}
}

func (t *LogMidBuilder) SaveReq() *LogMidBuilder {
	t.saveReq = true
	return t
}

func (t *LogMidBuilder) SaveResp() *LogMidBuilder {
	t.saveResp = true
	return t
}

func (t *LogMidBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var logData = LogData{}

		start := time.Now()
		logData.ReqTime = start

		method := ctx.Request.Method
		logData.Method = method

		path := ctx.Request.URL.Path
		logData.RespData = path

		header := map[string][]string{}
		for k, v := range ctx.Request.Header {
			for _, v2 := range v {
				header[k] = append(header[k], v2)
			}
		}
		logData.Header = header

		if t.saveReq {
			// Request.Body 是一个流，只能读一次。需要先读出来，然后再放回去。
			body, _ := ctx.GetRawData()
			logData.ReqData = string(body)
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
		}

		if t.saveResp {
			ctx.Writer = &responseWriter{
				ResponseWriter: ctx.Writer,
				data:           &logData,
			}
		}

		defer func() {
			logData.TimeCost = time.Since(start)
			t.saveLog(ctx, logData)
		}()

		ctx.Next()

		end := time.Now()
		logData.ReqTime = end
	}
}

// 装饰器思路，让 gin.Context.Writer 写响应数据的同时，再存一份日志
type responseWriter struct {
	gin.ResponseWriter
	data *LogData
}

func (t *responseWriter) WriteHeader(statusCode int) {
	t.data.StatusCode = statusCode
	t.ResponseWriter.WriteHeader(statusCode)
}

func (t *responseWriter) Write(data []byte) (int, error) {
	t.data.RespData = string(data)
	return t.ResponseWriter.Write(data)
}
