package context

import ctrl "sigs.k8s.io/controller-runtime"

var requestContext = "_context_request"

// Request returns the current operator request from the context
func Request(ctx Context) ctrl.Request {
	return ctx.Value(&requestContext).(ctrl.Request)
}

// WithRequest inserts the current operator request into the context
func WithRequest(ctx Context, request ctrl.Request) Context {
	return WithValue(ctx, &requestContext, request)
}
