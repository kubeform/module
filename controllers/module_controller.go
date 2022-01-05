/*
Copyright AppsCode Inc. and Contributors

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

package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	tfv1alpha1 "github.com/shahincsejnu/module-controller/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	meta_util "kmodules.xyz/client-go/meta"
	"path/filepath"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

const KFCFinalizer = "kfc.io"

var (
	basePath = filepath.Join("/tmp", ".kfc")
)

// ModuleReconciler reconciles a Module object
type ModuleReconciler struct {
	client.Client
	Log    logr.Logger
	Gvk    schema.GroupVersionKind
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=tf.kubeform.com,resources=modules,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tf.kubeform.com,resources=modules/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tf.kubeform.com,resources=modules/finalizers,verbs=update

func (r *ModuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//_ = log.FromContext(ctx)
	fmt.Println("got into the Reconcile")
	log := r.Log.WithValues("module", req.NamespacedName)
	fmt.Println("before gvk")
	// TODO(user): your logic here
	gvk := r.Gvk
	var obj unstructured.Unstructured
	obj.SetGroupVersionKind(gvk)
	fmt.Println("before the r.Get")
	if err := r.Get(ctx, req.NamespacedName, &obj); err != nil {
		log.Error(err, "unable to fetch Module")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	rClient := r.Client
	fmt.Println("before StartProcess")
	return ctrl.Result{}, StartProcess(rClient, ctx, gvk.GroupVersion(), &obj)
}

// SetupWithManager sets up the controller with the Manager.
func (r *ModuleReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tfv1alpha1.Module{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return !meta_util.MustAlreadyReconciled(e.Object)
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return (e.ObjectNew.(metav1.Object)).GetDeletionTimestamp() != nil || !meta_util.MustAlreadyReconciled(e.ObjectNew)
			},
		}).
		Complete(r)
}
