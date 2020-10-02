package kubernetes

import (
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/pkg/errors"

	containersv1alpha1 "github.com/Kafei59/bol/api/v1alpha1"
	"github.com/Kafei59/bol/pkg/context"
)

// ListDatabaseAuthorizations fetches all db auth resources from kubernetes which fits option list
func ListDatabaseAuthorizations(ctx context.Context, opts *client.ListOptions) (*containersv1alpha1.DatabaseAuthorizationList, error) {
	kubeClient := context.KubernetesClient(ctx)

	auths := &containersv1alpha1.DatabaseAuthorizationList{}
	err := kubeClient.List(ctx, auths, opts)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to list database authorizations")
	}

	return auths, nil
}

// GetDatabaseAuthorization fetches a db auth resource from kubernetes
func GetDatabaseAuthorization(ctx context.Context, namespace string, name string) (*containersv1alpha1.DatabaseAuthorization, error) {
	kubeClient := context.KubernetesClient(ctx)
	auth := &containersv1alpha1.DatabaseAuthorization{}

	err := kubeClient.Get(ctx, types.NamespacedName{Namespace: namespace, Name: name}, auth)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get database authorization resource")
	}

	// If UID is not set, returns an empty pointer
	if auth.UID == "" {
		return nil, nil
	}

	return auth, nil
}

// UpdateDatabaseAuthorization updates a db auth metadata fields
func UpdateNodePool(ctx context.Context, auth *containersv1alpha1.DatabaseAuthorization) error {
	kubeClient := context.KubernetesClient(ctx)

	err := kubeClient.Update(ctx, auth)
	if err != nil {
		return errors.Wrap(err, "Unable to update database authorization resource")
	}

	return nil
}

// UpdateDatabaseAuthorizationStatus update a db auth statuses
func UpdateDatabaseAuthorizationStatus(ctx context.Context, auth *containersv1alpha1.DatabaseAuthorization) error {
	kubeClient := context.KubernetesClient(ctx)

	err := kubeClient.Status().Update(ctx, auth)
	if err != nil {
		return errors.Wrap(err, "Unable to update database authorization resource statuses")
	}

	return nil
}

// DeleteDatabaseAuthorization removes a db auth resource from kubernetes
func DeleteDatabaseAuthorization(ctx context.Context) error {
	kubeClient := context.KubernetesClient(ctx)
	auth := context.DatabaseAuthorization(ctx)

	err := kubeClient.Delete(ctx, auth, nil)
	if err != nil {
		return errors.Wrap(err, "Unable to delete database authorization resource")
	}

	return nil
}
