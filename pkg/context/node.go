package context

import "k8s.io/api/core/v1"

var nodeContext = "_context_node"

// Node returns a kubernetes node resource from the context
func Node(ctx Context) *v1.Node {
	return ctx.Value(&nodeContext).(*v1.Node)
}

// WithNode inserts a kubernetes node resource into the context
func WithNode(ctx Context, node *v1.Node) Context {
	return WithValue(ctx, &nodeContext, node)
}
