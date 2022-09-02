package data

import (
	"context"
	"github.com/mittacy/go-toy-layout/apierr"
	"github.com/mittacy/go-toy-layout/app/model"
	"github.com/spf13/viper"
	"time"
)

type User struct {
	//eMysql.EGorm
	//eMongo.EMongo
	//eRedis.ERedis
	//client *thirdHttp.Client
}

func NewUser() User {
	return User{
		//EGorm:  eMysql.EGorm{ConfName: "localhost"},
		//EMongo: eMongo.EMongo{ConfName: "localhost", Collection: "collection_name"},
		//ERedis: eRedis.ERedis{ConfName: "localhost", DB: 0},
		//client: thirdHttp.NewClient(viper.GetString("user_server_host")),
	}
}

func (ctl *User) Get(c context.Context, id int64) (*model.User, error) {
	if id != 1 {
		return nil, apierr.UserNoExists
	}

	return &model.User{
		Id:        id,
		Name:      "xiyangyang",
		Age:       5,
		IsDeleted: model.UserIsDeletedNo,
		CreatedAt: time.Date(2022, 7, 1, 21, 1, 1, 0, time.Local),
		UpdatedAt: time.Date(2022, 7, 3, 12, 2, 15, 0, time.Local),
	}, nil
}

/*
 * 以下为查询缓存KEY方法
 */
func (ctl *User) cacheKeyPre() string {
	return viper.GetString("APP_NAME") + ":user"
}
