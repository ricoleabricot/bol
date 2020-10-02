package ovh

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/Kafei59/bol/pkg/context"
)

// PrivateDatabaseTask defines the structure of a task for private db management
type PrivateDatabaseTask struct {
	ID     int        `json:"id"`
	Status TaskStatus `json:"status"`
}

// GetPrivateDatabaseAuthorizedList returns a list of ips from an authorized list of a service
func GetPrivateDatabaseAuthorizedList(ctx context.Context, serviceName string) ([]string, error) {
	client := context.OvhClient(ctx)
	ips := make([]string, 0)

	// Request GET endpoint
	url := fmt.Sprintf("/hosting/privateDatabase/%s/whitelist?service=true", serviceName)
	err := client.Get(url, &ips)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get ip list from authorized list")
	}

	return ips, nil
}

// PostPrivateDatabaseAuthorizedList adds an ip to an authorized list for a service
func PostPrivateDatabaseAuthorizedList(ctx context.Context, serviceName string, ip string) error {
	client := context.OvhClient(ctx)
	task := &PrivateDatabaseTask{}

	// Create request body
	reqBody := map[string]string{
		"ip":      ip,
		"service": "true",
	}

	// Request POST endpoint
	url := fmt.Sprintf("/hosting/privateDatabase/%s/whitelist", serviceName)
	err := client.Post(url, reqBody, task)
	if err != nil {
		return errors.Wrap(err, "Unable to post ip to authorized list")
	}

	// Check task has been properly created
	if task.Status != TaskTodo {
		return errors.Errorf("Task has not been properly created, task status: %s", task.Status)
	}

	return WaitPrivateDatabaseTask(ctx, serviceName, task)
}

// DeletePrivateDatabaseAuthorizedList removes an ip from an authorized list of a service
func DeletePrivateDatabaseAuthorizedList(ctx context.Context, serviceName string, ip string) error {
	client := context.OvhClient(ctx)
	task := &PrivateDatabaseTask{}

	// Request DELETE endpoint
	url := fmt.Sprintf("/hosting/privateDatabase/%s/whitelist/%s", serviceName, ip)
	err := client.Delete(url, task)
	if err != nil {
		return errors.Wrap(err, "Unable to delete ip from authorized list")
	}

	// Check task has been properly created
	if task.Status != TaskTodo {
		return errors.Errorf("Task has not been properly created, task status: %s", task.Status)
	}

	return WaitPrivateDatabaseTask(ctx, serviceName, task)
}

// WaitPrivateDatabaseTask waits for privateDatabase task be removed or marked as done (or any error)
func WaitPrivateDatabaseTask(ctx context.Context, serviceName string, task *PrivateDatabaseTask) error {
	client := context.OvhClient(ctx)

	url := fmt.Sprintf("/hosting/privateDatabase/%s/tasks/%d", serviceName, task.ID)

	// Wait 10 times to see if the task have been removed or any error happened
	for i := 0; i < 10; i++ {
		fetchedTask := &PrivateDatabaseTask{}
		err := client.Get(url, &fetchedTask)

		// Check if the task have been deleted
		if err != nil {
			if IsNotFoundError(err) {
				return nil
			}

			return err
		}

		// Otherwise, check task statuses
		switch task.Status {
		case TaskDone:
			return nil
		case TaskError:
			return errors.Errorf("Task %d is in error status", task.ID)
		case TaskCancelled:
			return errors.Errorf("Task %d have been cancelled", task.ID)
		}

		// If it's still in to do, just wait and check later
		time.Sleep(2 * time.Second)
	}

	return errors.Errorf("Task %d is still in %s status after 10 requests, abort", task.ID, task.Status)
}
