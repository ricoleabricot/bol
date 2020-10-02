/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package node

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"

	"github.com/Kafei59/bol/controllers"
	"github.com/Kafei59/bol/pkg/context"
	"github.com/Kafei59/bol/pkg/kubernetes"
)

// NodeReconciler reconciles a Node object
type NodeReconciler struct {
	client.Client

	Log    logr.Logger
	Scheme *runtime.Scheme
}

func (r *NodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Node{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=,resources=nodes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=,resources=nodes/status,verbs=get;update;patch

func (r *NodeReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := controllers.InitContext(r, req, r.Log)

	// Get the node resource from request
	node, err := kubernetes.GetNode(ctx, req.Namespace, req.Name)
	if err != nil {
		return controllers.ResourceNotFoundError(ctx, err)
	}

	// If node resource is set to deleting, clean and remove everything properly
	if !node.GetDeletionTimestamp().IsZero() {
		return r.DeleteResource(ctx)
	}

	// Otherwise, it should update the resource idempotently
	return r.UpdateResource(ctx)
}

func (r *NodeReconciler) UpdateResource(ctx context.Context) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (r *NodeReconciler) DeleteResource(ctx context.Context) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}
