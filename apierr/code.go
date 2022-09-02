package apierr

import (
	"github.com/mittacy/go-toy/core/bizerr"
)

// 业务错误定义格式: 系统码:模块码:错误相关码
var (
	// 用户相关code: 1101XX
	UserNoExists = &bizerr.BizErr{Code: 110101, Msg: "用户不存在"}
)
