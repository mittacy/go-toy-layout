package service

import (
	"context"
	"github.com/mittacy/go-toy-layout/app/data"
	"github.com/mittacy/go-toy-layout/app/service/smodel"
	"github.com/mittacy/go-toy/core/singleton"
	"github.com/pkg/errors"
)

// 一般情况下service应该只引用并控制自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型

// User 服务说明注释
var User userService

type userService struct {
	data data.User
}

func init() {
	singleton.Register(func() {
		User = userService{
			data: data.NewUser(),
		}
	})
}

func (ctl *userService) GetById(c context.Context, id int64) (*smodel.GetById, error) {
	user, err := ctl.data.Get(c, id)
	if err != nil {
		return nil, errors.WithMessage(err, "查询记录错误")
	}

	res := &smodel.GetById{User: user}
	return res, nil
}
