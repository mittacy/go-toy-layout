package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/go-toy-layout/bootstrap"
	"github.com/mittacy/go-toy-layout/config"
	"github.com/mittacy/go-toy-layout/middleware"
	"github.com/mittacy/go-toy-layout/router"
	"github.com/mittacy/go-toy/core/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

var StartCmd = &cobra.Command{
	Use:   "http",
	Short: "start the http api server. Example: go run . start http",
	Long:  "start the http api server. Example: go run . start http -c=.env.development -e=development -p=8080",
	Run:   run,
}

var (
	port int
	conf string
	env  string
)

func init() {
	StartCmd.Flags().StringVarP(&conf, "conf", "c", ".env.development", "conf file path, example: .env.production")
	StartCmd.Flags().StringVarP(&env, "env", "e", "development", "running environment, example: development, optional: development/test/production")
	StartCmd.Flags().IntVarP(&port, "port", "p", 8080, "listen port, example: 8080")
}

func run(cmd *cobra.Command, args []string) {
	fmt.Printf("conf: %s, port: %d, env: %s\n", conf, port, env)
	bootstrap.InitDependents(conf, env)

	r := gin.New()
	// 全局中间件
	r.Use(middleware.Recovery())
	config.InitZipkinTracer(r)

	router.Init(r)

	if err := serve(r); err != nil {
		log.Panicf("http panic, err: %+v", err)
	}
}

// serve 启动服务
func serve(r *gin.Engine) error {
	readTimeout := viper.GetDuration("APP_READ_TIMEOUT")
	if readTimeout == 0 {
		readTimeout = 5
	}
	writeTimeout := viper.GetDuration("APP_WRITE_TIMEOUT")
	if writeTimeout == 0 {
		writeTimeout = 5
	}

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        r,
		ReadTimeout:    time.Second * readTimeout,
		WriteTimeout:   time.Second * writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
