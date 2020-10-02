package context

import "github.com/go-logr/logr"

var loggerContext = "_context_logger"

// Logger returns the logger from the context
func Logger(ctx Context) logr.Logger {
	return ctx.Value(&loggerContext).(logr.Logger)
}

// WithLogger inserts a logger into the context
func WithLogger(ctx Context, logger logr.Logger) Context {
	return WithValue(ctx, &loggerContext, logger)
}
