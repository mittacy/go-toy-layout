package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/go-toy-layout/config"
	"github.com/mittacy/go-toy-layout/utils/bizUtil"
	"github.com/mittacy/go-toy-layout/variable"
	"github.com/mittacy/go-toy/core/response"
	"github.com/mittacy/go-toy/core/singleton"
)

// InitDependents 初始化依赖资源
// @param path 配置文件路径
// @param port 端口
// @param env 运行环境
func InitDependents(path string, env string) {
	// conf
	ParseConfig("env", path)

	// gin run env
	gin.SetMode(bizUtil.AppEnvToGinEnv(env))

	// log
	InitLog()

	// go http
	InitGoHttp()

	// database
	config.InitMysql()
	config.InitRedis()
	config.InitMongo()

	// 响应体增加自定义字段
	response.WithFieldKeysFromCtx(variable.TraceID, variable.SpanID)

	// 初始化单例
	singleton.Build()
}
