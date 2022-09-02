package dp

import (
	"context"
	"github.com/mittacy/go-toy-layout/app/service/smodel"
	"github.com/mittacy/go-toy-layout/app/validator/userVdr"
	"github.com/mittacy/go-toy-layout/utils/timeUtil"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (ctl *User) Get(c context.Context, data *smodel.GetById) map[string]interface{} {
	user := userVdr.GetReplyUser{
		Id:        data.User.Id,
		Name:      data.User.Name,
		Age:       data.User.Age,
		CreatedAt: timeUtil.Format(data.User.CreatedAt),
		UpdatedAt: timeUtil.Format(data.User.UpdatedAt),
	}

	return map[string]interface{}{
		"user": user,
	}
}
