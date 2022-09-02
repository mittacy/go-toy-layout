package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/go-toy/core/log"
	"github.com/mittacy/go-toy/core/response"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recoverErr := recover(); recoverErr != nil {
				err := errors.New(fmt.Sprintf("%s", recoverErr))
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				log.ErrorwWithCtx(c, "Recovery from panic", "request", string(httpRequest), "err", err)
				response.CustomErr(c, http.StatusOK, http.StatusInternalServerError, err, response.NilData)
				return
			}
		}()
		c.Next()
	}
}
