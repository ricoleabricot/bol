package kubernetes

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/pkg/errors"

	"github.com/Kafei59/bol/pkg/context"
)

// ListNodes fetches all nodes resources from kubernetes which fits option list
func ListNodes(ctx context.Context, opts *client.ListOptions) (*v1.NodeList, error) {
	kubeClient := context.KubernetesClient(ctx)

	nodes := &v1.NodeList{}
	err := kubeClient.List(ctx, nodes, opts)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to list nodes")
	}

	return nodes, nil
}

// GetNode fetches a node resource from kubernetes
func GetNode(ctx context.Context, namespace string, name string) (*v1.Node, error) {
	kubeClient := context.KubernetesClient(ctx)
	node := &v1.Node{}

	err := kubeClient.Get(ctx, types.NamespacedName{Namespace: namespace, Name: name}, node)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get node resource")
	}

	// If UID is not set, returns an empty pointer
	if node.UID == "" {
		return nil, nil
	}

	return node, nil
}

// GetNodeIP computes the address of the node into a CIDR
func GetNodeIP(ctx context.Context, node *v1.Node) (string, error) {
	for _, address := range node.Status.Addresses {
		switch address.Type {
		case v1.NodeInternalIP:
			return fmt.Sprintf("%s/32", address.Address), nil
		default:
			// Do nothing
		}
	}

	return "", errors.New("unable to get node internal IP")
}
