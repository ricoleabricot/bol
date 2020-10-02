package v1alpha1

import "fmt"

// String returns a readable string of a resource
func (da *DatabaseAuthorization) String() string {
	return fmt.Sprintf(
		"%s:%s/%s",
		da.GetObjectKind().GroupVersionKind().GroupVersion().Version, da.GetNamespace(), da.Name,
	)
}

// HasBeenUpdated checks if a resource has a new generation metadata (it is increments in each update except for meta and status)
func (da *DatabaseAuthorization) HasBeenUpdated() bool {
	oldGeneration := da.Status.ObservedGeneration
	newGeneration := da.ObjectMeta.Generation

	return oldGeneration != 0 && oldGeneration != newGeneration
}

// IsDeleting checks if a resource is set to be deleted
func (da *DatabaseAuthorization) IsDeleting() bool {
	return !da.GetObjectMeta().GetDeletionTimestamp().IsZero()
}

// GetFinalizers fetches all finalizers bound to a resource
func (da *DatabaseAuthorization) GetFinalizers() []string {
	return da.GetObjectMeta().GetFinalizers()
}

// SetFinalizers sets all finalizers on a resource
func (da *DatabaseAuthorization) SetFinalizers(finalizers []string) {
	da.GetObjectMeta().SetFinalizers(finalizers)
}

// HasFinalizer checks if the resource is bound to a finalizer
func (da *DatabaseAuthorization) HasFinalizer(finalizer string) bool {
	finalizers := da.GetFinalizers()

	for _, f := range finalizers {
		if f == finalizer {
			return true
		}
	}

	return false
}

// AddFinalizer binds a finalizer to the resource if needed
func (da *DatabaseAuthorization) AddFinalizer(finalizer string) bool {
	if da.HasFinalizer(finalizer) {
		return false
	}

	finalizers := append(da.GetFinalizers(), finalizer)
	da.SetFinalizers(finalizers)

	return true
}

// RemoveFinalizer removes the resource finalizer if it's bound
func (da *DatabaseAuthorization) RemoveFinalizer(finalizer string) bool {
	if !da.HasFinalizer(finalizer) {
		return false
	}

	finalizers := da.GetFinalizers()
	for i, value := range finalizers {
		if value == finalizer {
			finalizers = append(finalizers[:i], finalizers[i+1:]...)
			break
		}
	}

	da.SetFinalizers(finalizers)

	return true
}
