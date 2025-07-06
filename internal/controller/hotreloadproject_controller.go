/*
Copyright 2025 Akos Orban.

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

package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"akosrbn.io/hot-reload/api/v1alpha1"
	reloaderv1alpha1 "akosrbn.io/hot-reload/api/v1alpha1"
)

// lastReloadAnnotation is the annotation key that stores the timestamp
// of the last hot-reload operation in RFC3339 format.
const lastReloadAnnotation = "reloader.akosrbn.io/rebuilding"

type realClock struct{}

func (_ realClock) Now() time.Time { return time.Now() }

// Clock knows how to get the current time.
// It can be used to fake out timing for testing.
type Clock interface {
	Now() time.Time
}

// HotReloadProjectReconciler reconciles a HotReloadProject object
type HotReloadProjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Clock
}

// For HotReloadProject resources
// +kubebuilder:rbac:groups=reloader.akosrbn.io,resources=hotreloadprojects,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=reloader.akosrbn.io,resources=hotreloadprojects/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=reloader.akosrbn.io,resources=hotreloadprojects/finalizers,verbs=update

// For Deployments (to restart/update containers)
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HotReloadProject object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *HotReloadProjectReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	var project v1alpha1.HotReloadProject
	if err := r.Get(ctx, req.NamespacedName, &project); err != nil {
		log.Error(err, "Unable to fetch HotReloadProject")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HotReloadProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&reloaderv1alpha1.HotReloadProject{}).
		Named("hotreloadproject").
		Complete(r)
}
