package context

import "github.com/ovh/go-ovh/ovh"

var ovhContext = "_context_ovh_client"

// OvhClient returns an OVHcloud client from the context
func OvhClient(ctx Context) *ovh.Client {
	return ctx.Value(&ovhContext).(*ovh.Client)
}

// WithOvhClient inserts an OVHcloud client in the context
func WithOvhClient(ctx Context, client *ovh.Client) Context {
	return WithValue(ctx, &ovhContext, client)
}
