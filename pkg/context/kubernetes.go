package context

import "sigs.k8s.io/controller-runtime/pkg/client"

var kubeContext = "_context_kubernetes_client"

// KubernetesClient returns a kubernetes client from the context
func KubernetesClient(ctx Context) client.Client {
	return ctx.Value(&kubeContext).(client.Client)
}

// WithKubernetesClient inserts a kubernetes client in the context
func WithKubernetesClient(ctx Context, client client.Client) Context {
	return WithValue(ctx, &kubeContext, client)
}
