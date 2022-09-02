package apierr

import "github.com/mittacy/go-toy/core/bizerr"

func New(code int, msg string) *bizerr.BizErr {
	return &bizerr.BizErr{
		Code: code,
		Msg:  msg,
	}
}
