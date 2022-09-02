package config

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/go-toy-layout/variable"
	"github.com/mittacy/go-toy/core/log"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/openzipkin/zipkin-go/propagation/b3"
	"github.com/openzipkin/zipkin-go/reporter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"strconv"
)

var ZKReporter reporter.Reporter

// 初始化zipkin客户端，并将服务注册到zipkin
func InitZipkinTracer(engine *gin.Engine) {
	// set up a span reporter
	ZKReporter = newReporter("request")

	// initialize our tracer
	tracer, err := zipkin.NewTracer(ZKReporter)
	if err != nil {
		log.Panicw("init zipkin trace err, unable to create tracer", "err", err)
	}

	engine.Use(func(c *gin.Context) {
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		var bodyBytes []byte
		if c.ContentType() == gin.MIMEJSON {
			var err error
			bodyBytes, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.ErrorwWithCtx(c, "Invalid request body", "err", err)
				c.Abort()
				return
			}

			// 新建缓冲区并替换原有Request.body
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		reqData := map[string]string{
			"method":     c.Request.Method,
			"path":       path,
			"query":      query,
			"ip":         c.ClientIP(),
			"user-agent": c.Request.UserAgent(),
			"body":       string(bodyBytes),
		}

		// 解析父请求信息
		extractor := b3.ExtractHTTP(c.Request)
		parentSpan, _ := extractor()
		spanOptions := []zipkin.SpanOption{zipkin.Tags(reqData)}
		if parentSpan != nil {
			spanOptions = append(spanOptions, zipkin.Parent(*parentSpan))
		}
		span := tracer.StartSpan(c.FullPath(), spanOptions...)

		// 写入上下文
		c.Set(variable.TraceID, span.Context().TraceID.String())
		c.Set(variable.SpanID, span.Context().ID.String())
		c.Set(variable.Sampled, span.Context().Sampled)
		c.Set(variable.SpanCtxKey, span.Context())
		defer span.Finish()

		c.Next()

		// 记录响应数据
		span.Tag("status", strconv.Itoa(c.Writer.Status()))
		span.Tag("resp", blw.body.String())
	})
}

type ZKLog struct {
	l *log.Logger
}

func newReporter(logName string) reporter.Reporter {
	return &ZKLog{
		l: log.New(logName),
	}
}

// Send outputs a span to the Go logger.
// 将信息存入上下文
func (z *ZKLog) Send(s model.SpanModel) {
	fields := []zapcore.Field{
		zap.String("status", s.Tags["status"]),
		zap.String("method", s.Tags["method"]),
		zap.String("path", s.Tags["path"]),
		zap.String("query", s.Tags["query"]),
		zap.String("ip", s.Tags["ip"]),
		zap.String("user-agent", s.Tags["user-agent"]),
		zap.String("body", s.Tags["body"]),
		zap.String("resp", s.Tags["resp"]),
		zap.String("X-B3-Flags", "0"),
		zap.String("X-B3-SpanId", s.SpanContext.ID.String()),
		zap.String("X-B3-TraceId", s.TraceID.String()),
		zap.Duration("X-B3-Duration", s.Duration),
		zap.Time("X-B3-Timestamp", s.Timestamp),
	}

	parentSpanId := ""
	if s.SpanContext.ParentID != nil {
		parentSpanId = s.ParentID.String()
	}
	fields = append(fields, zap.String("X-B3-ParentSpanId", parentSpanId))

	if s.Sampled != nil {
		fields = append(fields, zap.Bool("X-B3-Sampled", *s.Sampled))
	}

	z.l.Info(s.Name, fields...)
}

// Close closes the reporter
func (z *ZKLog) Close() error {
	return z.l.Sync()
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
