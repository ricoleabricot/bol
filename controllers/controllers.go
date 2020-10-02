package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"

	"github.com/Kafei59/bol/pkg/context"
	"github.com/Kafei59/bol/pkg/kubernetes"
)

// InitContext sets all default values to the controller context
func InitContext(cli client.Client, req ctrl.Request, log logr.Logger) context.Context {
	// Forge background context
	ctx := context.Background()

	// Set all useful data to context
	ctx = context.WithKubernetesClient(ctx, cli)
	ctx = context.WithRequest(ctx, req)
	ctx = context.WithLogger(ctx, log)

	return ctx
}

// ResourceNotFoundError checks if the resource currently watched has not been already removed
func ResourceNotFoundError(ctx context.Context, err error) (ctrl.Result, error) {
	req := context.Request(ctx)
	log := context.Logger(ctx)

	// If the resource has already been deleted, do not ask for reconciliation requeue
	if kubernetes.IsNotFoundErr(err) {
		log.Info("This resource does not exist anymore, do not watch it anymore", "name", req.Name)

		return ctrl.Result{}, nil
	}

	log.Error(err, "Unable to get the resource", "name", req)

	return ctrl.Result{Requeue: true}, err
}
