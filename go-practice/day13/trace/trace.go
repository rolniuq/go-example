package trace

import "context"

type ctxKey string

const TraceIDKey ctxKey = "trace-id"

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(TraceIDKey).(string)
	if !ok {
		return "unknown"
	}

	return traceID
}
