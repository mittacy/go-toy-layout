package bootstrap

import (
	"github.com/mittacy/go-toy-layout/utils/timeUtil"
	"github.com/mittacy/go-toy-layout/variable"
	"github.com/mittacy/go-toy/core/log"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func InitLog() {
	logPath := viper.GetString("LOG_PATH")
	logLevel := viper.GetString("LOG_LOW_LEVEL")
	logEncoderJson := viper.GetBool("LOG_ENCODER_JSON")
	logInConsole := viper.GetBool("Log_In_Console")
	globalFields := []zapcore.Field{
		{
			Key:    "module_name",
			Type:   zapcore.StringType,
			String: viper.GetString("APP_NAME"),
		},
	}

	log.SetDefaultConf(
		log.WithPath(logPath),
		log.WithTimeFormat(timeUtil.TimeFormat),
		log.WithLevel(logLevel),
		log.WithEncoderJSON(logEncoderJson),
		log.WithFields(globalFields...),
		log.WithCtxField(variable.TraceID, variable.SpanID),
		log.WithLogInConsole(logInConsole))
}
