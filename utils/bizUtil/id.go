package bizUtil

import (
	"fmt"
	"github.com/mittacy/go-toy-layout/utils/randomUtil"
	"time"
)

// NewRequestId 生成新的请求id
func NewRequestId() string {
	now := time.Now()
	s := fmt.Sprintf("%s%08x%05x", "r", now.Unix(), now.UnixNano()%0x100000)
	return s + "_" + randomUtil.String(18)
}
