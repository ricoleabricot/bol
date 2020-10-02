package authorization

import (
	"k8s.io/apimachinery/pkg/labels"
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
	ctx = context.WithDatabaseAuthorization(ctx, auth)

	// If db auth resource is set to deleting, clean and remove everything properly
	if auth.IsDeleting() {
		return r.DeleteResource(ctx)
	}

	// Otherwise, it should update the resource idempotently
	return r.UpdateResource(ctx)
}

// UpdateResource will apply all modifications on database authorized lists
func (r *DatabaseAuthorizationReconciler) UpdateResource(ctx context.Context) (ctrl.Result, error) {
	auth := context.DatabaseAuthorization(ctx)
	log := context.Logger(ctx)

	log.WithValues("event", "update")

	// Attach finalizer if needed
	done := auth.AddFinalizer()
	if done {
		log.Info("Set finalizer to resource")
	}

	// Forge list options with authorization labels selector
	opts := &client.ListOptions{}
	if len(auth.Spec.LabelSelector.MatchLabels) > 0 {
		sel := labels.SelectorFromSet(auth.Spec.LabelSelector.MatchLabels)
		opts.LabelSelector = sel
	}

	// Fetch all nodes resources which fits authorization label selectors
	nodes, err := kubernetes.ListNodes(ctx, opts)
	if err != nil {
		log.Error(err, "Unable to list nodes")

		return ctrl.Result{}, err
	}

	// TODO: set in status all authorized ips to be able to remove it afterwards
	// TODO: find a way to be able to remove authorized list when label / services changes in auth resource

	for _, node := range nodes.Items {
		ip, err := kubernetes.GetNodeIP(ctx, &node)
		if err != nil {
			log.Error(err, "Unable to get node ip")

			return ctrl.Result{}, err
		}

		log.Info("IP should be authorized", "ip", ip)
	}

	return ctrl.Result{}, nil
}

// DeleteResource will remove all database authorized lists
func (r *DatabaseAuthorizationReconciler) DeleteResource(ctx context.Context) (ctrl.Result, error) {
	_ = context.DatabaseAuthorization(ctx)
	log := context.Logger(ctx)

	log.WithValues("event", "update")

	return ctrl.Result{}, nil
}
