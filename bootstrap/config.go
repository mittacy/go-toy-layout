package bootstrap

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// ParseConfig 解析
func ParseConfig(configType string, configPath string) {
	viper.SetConfigType(configType)
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("load config file fail: %s", err)
	}

	// 监听配置实时更新
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.SetConfigFile(e.Name)
		log.Printf("some configuration item in the %s file has changed", e.Name)
	})
	viper.WatchConfig()
}
