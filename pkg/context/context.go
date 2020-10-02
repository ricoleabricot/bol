package context

import "context"

type Context context.Context

// Background starts a new context
func Background() Context {
	return context.Background()
}

// WithValue adds a new key/val pair in the current context
func WithValue(ctx Context, key interface{}, val interface{}) Context {
	return context.WithValue(ctx, key, val)
}
