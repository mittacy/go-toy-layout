package bootstrap

import (
	"github.com/mittacy/go-toy-layout/variable"
	"github.com/mittacy/go-toy/core/goHttp"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func InitGoHttp() {
	// 初始化api请求客户端
	maxIdleConnsPerHost := viper.GetInt("GoHTTPMaxIdleConnsPerHost")
	if maxIdleConnsPerHost <= 0 {
		maxIdleConnsPerHost = 1000
	}
	MaxConnsPerHost := viper.GetInt("GoHTTPMaxConnsPerHost")
	if maxIdleConnsPerHost <= 0 {
		MaxConnsPerHost = 1000
	}
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   maxIdleConnsPerHost, // 单个Host的最大空闲连接数
			MaxConnsPerHost:       MaxConnsPerHost,     // 单个Host的最大连接总数，>=MaxIdleConnsPerHost
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: time.Second * 5,
	}
	goHttp.Init(client, "thirdHttp", goHttp.WithHeaderB3TraceFromCtx(variable.SpanCtxKey))
}
