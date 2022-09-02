package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/go-toy-layout/app/dp"
	"github.com/mittacy/go-toy-layout/app/service"
	"github.com/mittacy/go-toy-layout/app/validator/userVdr"
	"github.com/mittacy/go-toy/core/log"
	"github.com/mittacy/go-toy/core/response"
	"github.com/mittacy/go-toy/core/singleton"
)

var User userApi

type userApi struct {
	dp dp.User
}

func init() {
	singleton.Register(func() {
		User = userApi{
			dp: dp.NewUser(),
		}
	})
}

func (ctl *userApi) Get(c *gin.Context) {
	req := userVdr.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}

	data, err := service.User.GetById(c, req.Id)
	if err != nil {
		response.FailCheckBizErr(c, log.Default(), req, "查询记录错误", err)
		return
	}

	res := ctl.dp.Get(c, data)
	response.Success(c, res)
}
