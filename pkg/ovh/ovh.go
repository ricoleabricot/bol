package ovh

import (
	"net/http"

	v1 "k8s.io/api/core/v1"

	"github.com/ovh/go-ovh/ovh"
)

// TaskStatus exposes an enum for status type of a task
type TaskStatus string

const (
	TaskTodo      TaskStatus = "todo"
	TaskDone      TaskStatus = "done"
	TaskError     TaskStatus = "error"
	TaskCancelled TaskStatus = "cancelled"
)

// CreateOvhClient creates an OVHcloud client to call endpoints for database authorization
func CreateOvhClient(appSecret *v1.Secret, tokenSecret *v1.Secret) (*ovh.Client, error) {
	// https://github.com/ovh/go-ovh#configuration
	return ovh.NewClient(
		string(appSecret.Data["endpoint"]),
		string(appSecret.Data["key"]),
		string(appSecret.Data["secret"]),
		string(tokenSecret.Data["token"]),
	)
}

// IsNotFoundError checks if an error has a 404 HTTP Code
func IsNotFoundError(err error) bool {
	apiErr, ok := err.(*ovh.APIError)
	if ok {
		return apiErr.Code == http.StatusNotFound
	}

	return false
}

// IsConflictError checks if an error has a 409 HTTP Code
func IsConflictError(err error) bool {
	apiErr, ok := err.(*ovh.APIError)
	if ok {
		return apiErr.Code == http.StatusConflict
	}

	return false
}
