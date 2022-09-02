package bizUtil

import (
	"context"
	"github.com/mittacy/go-toy-layout/variable"
)

func GetTraceId(c context.Context) string {
	if v, ok := c.Value(variable.TraceID).(string); ok {
		return v
	}
	return ""
}
