package kubernetes

import (
	"github.com/pkg/errors"

	apiErrors "k8s.io/apimachinery/pkg/api/errors"
)

// IsNotFoundErr checks if the kubernetes resource does not exists
func IsNotFoundErr(err error) bool {
	return apiErrors.IsNotFound(errors.Cause(err))
}
