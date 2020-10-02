package kubernetes

import (
	"k8s.io/apimachinery/pkg/types"

	"github.com/pkg/errors"

	containersv1alpha1 "github.com/Kafei59/bol/api/v1alpha1"
	"github.com/Kafei59/bol/pkg/context"
)

// GetDatabaseAuthorization fetches a db auth resource from kubernetes
func GetDatabaseAuthorization(ctx context.Context, namespace string, name string) (*containersv1alpha1.DatabaseAuthorization, error) {
	client := context.KubernetesClient(ctx)
	auth := &containersv1alpha1.DatabaseAuthorization{}

	err := client.Get(ctx, types.NamespacedName{Namespace: namespace, Name: name}, auth)
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
	client := context.KubernetesClient(ctx)

	err := client.Update(ctx, auth)
	if err != nil {
		return errors.Wrap(err, "Unable to update database authorization resource")
	}

	return nil
}

// UpdateDatabaseAuthorizationStatus update a db auth statuses
func UpdateDatabaseAuthorizationStatus(ctx context.Context, auth *containersv1alpha1.DatabaseAuthorization) error {
	client := context.KubernetesClient(ctx)

	err := client.Status().Update(ctx, auth)
	if err != nil {
		return errors.Wrap(err, "Unable to update database authorization resource statuses")
	}

	return nil
}

// DeleteDatabaseAuthorization removes a db auth resource from kubernetes
func DeleteDatabaseAuthorization(ctx context.Context) error {
	client := context.KubernetesClient(ctx)
	auth := context.DatabaseAuthorization(ctx)

	err := client.Delete(ctx, auth, nil)
	if err != nil {
		return errors.Wrap(err, "Unable to delete database authorization resource")
	}

	return nil
}
