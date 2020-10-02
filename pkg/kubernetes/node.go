package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/pkg/errors"

	"github.com/Kafei59/bol/pkg/context"
)

// GetNode fetches a node resource from kubernetes
func GetNode(ctx context.Context, namespace string, name string) (*v1.Node, error) {
	client := context.KubernetesClient(ctx)
	node := &v1.Node{}

	err := client.Get(ctx, types.NamespacedName{Namespace: namespace, Name: name}, node)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get node resource")
	}

	// If UID is not set, returns an empty pointer
	if node.UID == "" {
		return nil, nil
	}

	return node, nil
}
