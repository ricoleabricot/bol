package authorization

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"

	containersv1alpha1 "github.com/Kafei59/bol/api/v1alpha1"
	"github.com/Kafei59/bol/controllers"
	"github.com/Kafei59/bol/pkg/context"
	"github.com/Kafei59/bol/pkg/kubernetes"
)

// DatabaseAuthorizationReconciler reconciles a DatabaseAuthorization object
type DatabaseAuthorizationReconciler struct {
	client.Client

	Log    logr.Logger
	Scheme *runtime.Scheme
}

func (r *DatabaseAuthorizationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&containersv1alpha1.DatabaseAuthorization{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=containers.ovhcloud.com,resources=databaseauthorizations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=containers.ovhcloud.com,resources=databaseauthorizations/status,verbs=get;update;patch

func (r *DatabaseAuthorizationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := controllers.InitContext(r, req, r.Log)

	// Get the authorization resource from request
	auth, err := kubernetes.GetDatabaseAuthorization(ctx, req.Namespace, req.Name)
	if err != nil {
		return controllers.ResourceNotFoundError(ctx, err)
	}

	switch true {
	// If db auth resource is set to deleting, clean and remove everything properly
	case auth.IsDeleting():
		return r.DeleteResource(ctx)

	// If we see any changes on the object generation observed, update the resource, otherwise, sync it
	case auth.HasBeenUpdated():
		return r.UpdateResource(ctx)

	// Otherwise, it should create the resource
	default:
		return r.CreateResource(ctx)
	}
}

func (r *DatabaseAuthorizationReconciler) CreateResource(ctx context.Context) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (r *DatabaseAuthorizationReconciler) UpdateResource(ctx context.Context) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (r *DatabaseAuthorizationReconciler) DeleteResource(ctx context.Context) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}
