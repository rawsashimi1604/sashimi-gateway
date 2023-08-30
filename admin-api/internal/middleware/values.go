package middleware

import "context"

// Defines a set of values that will be passed to the different middleware handlers.
type MiddlewareValues struct {
	StatusCode int
	Data       interface{}
	ServiceId  int
	RouteId    int
}

var MiddlewareValuesKey ContextKey = "middlewareValues"

func SetMiddlewareValues(ctx context.Context, values MiddlewareValues) context.Context {
	return context.WithValue(ctx, MiddlewareValuesKey, values)
}

func RetrieveMiddlewareValues(ctx context.Context) (MiddlewareValues, bool) {
	values, ok := ctx.Value(MiddlewareValuesKey).(MiddlewareValues)
	return values, ok
}
