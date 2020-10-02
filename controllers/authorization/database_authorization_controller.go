package authorization


import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	containersv1alpha1 "github.com/Kafei59/bol/api/v1alpha1"
)

// DatabaseAuthorizationReconciler reconciles a DatabaseAuthorization object
type DatabaseAuthorizationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=containers.ovhcloud.com,resources=databaseauthorizations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=containers.ovhcloud.com,resources=databaseauthorizations/status,verbs=get;update;patch

func (r *DatabaseAuthorizationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("databaseauthorization", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *DatabaseAuthorizationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&containersv1alpha1.DatabaseAuthorization{}).
		Complete(r)
}
