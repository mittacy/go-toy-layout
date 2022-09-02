package variable

import "github.com/openzipkin/zipkin-go/propagation/b3"

const (
	TraceID      = b3.TraceID
	SpanID       = b3.SpanID
	ParentSpanID = b3.ParentSpanID
	Sampled      = b3.Sampled
	Flags        = b3.Flags
	SpanCtxKey   = "b3SpanCtx"
)
