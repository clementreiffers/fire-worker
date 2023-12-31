/*
Copyright 2023 clementreiffers.

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
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1 "operators/WorkerBundle/api/v1"
)

// WorkerVersionReconciler reconciles a WorkerVersion object
type WorkerVersionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=api.cf-worker,resources=workerversions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=api.cf-worker,resources=workerversions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=api.cf-worker,resources=workerversions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the WorkerVersion object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile

func createWorkerRelease(instance *apiv1.WorkerVersion) apiv1.WorkerRelease {
	return apiv1.WorkerRelease{
		ObjectMeta: metav1.ObjectMeta{Name: getWorkerRelease(instance.Spec.Accounts), Namespace: instance.GetNamespace()},
		Spec: apiv1.WorkerReleaseSpec{
			WorkerVersions: map[string]string{
				instance.Spec.Scripts: instance.Spec.Url,
			},
			Accounts: instance.Spec.Accounts,
		},
	}
}

func (r *WorkerVersionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.Log.WithValues("WorkerVersion", req.NamespacedName)

	instance := &apiv1.WorkerVersion{}
	err := r.Get(ctx, req.NamespacedName, instance)

	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	workerRelease := apiv1.WorkerRelease{}
	err = r.Get(ctx, types.NamespacedName{Name: getWorkerRelease(instance.Spec.Accounts), Namespace: instance.GetNamespace()}, &workerRelease)
	if err != nil {
		workerRelease := createWorkerRelease(instance)
		err = r.Create(ctx, &workerRelease)
		if err != nil {
			return ctrl.Result{}, err
		}
		logger.Info("WorkerRelease created!")
		return ctrl.Result{}, nil

	} else {
		workerRelease.Spec.WorkerVersions[instance.Spec.Scripts] = instance.Spec.Url
		err = r.Update(ctx, &workerRelease)
		if err != nil {
			return ctrl.Result{}, err
		}
		logger.Info("WorkerRelease updated!")
		return ctrl.Result{}, nil
	}

}

// SetupWithManager sets up the controller with the Manager.
func (r *WorkerVersionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1.WorkerVersion{}).
		Complete(r)
}
