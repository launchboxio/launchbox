/*
Copyright 2022.

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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"reflect"
	"strconv"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	launchboxiov1alpha1 "github.com/launchboxio/launchbox/api/v1alpha1"
)

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=launchbox.io,resources=projects,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=launchbox.io,resources=projects/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=launchbox.io,resources=projects/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Project object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *ProjectReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	out := log.FromContext(ctx)

	project := &launchboxiov1alpha1.Project{}
	err := r.Get(ctx, req.NamespacedName, project)
	if err != nil {
		if errors.IsNotFound(err) {
			out.Info("Project resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		out.Error(err, "Failed to get Project")
		return ctrl.Result{}, err
	}
	found := &v1.ServiceAccount{}
	err = r.Get(ctx, types.NamespacedName{Name: project.Name, Namespace: project.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		sa := r.serviceAccountForProject(project)
		out.Info("Creating a new Service Account", "ServiceAccount.Namespace", project.Namespace, "ServiceAccount.Name", project.Name)
		err = r.Create(ctx, sa)
		if err != nil {
			out.Error(err, "Failed to create new service account", "ServiceAccount.Namespace", project.Namespace, "ServiceAccount.Name", project.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		out.Error(err, "Failed to get service account")
		return ctrl.Result{}, err
	}

	// TODO: Ensure labels and annotations are up to date

	// TODO: If prometheus.enabled, create the ServiceMonitor

	// Get the serviceAccount.name, and update the status
	if !reflect.DeepEqual(found.Name, project.Status.ServiceAccount) {
		project.Status.ServiceAccount = found.Name
		err := r.Status().Update(ctx, project)
		if err != nil {
			out.Error(err, "Failed to update Project status")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *ProjectReconciler) serviceAccountForProject(p *launchboxiov1alpha1.Project) *v1.ServiceAccount {
	labels := map[string]string{
		"launchbox.io/application.id": strconv.Itoa(int(p.Spec.ApplicationId)),
		"launchbox.io/project.id":     strconv.Itoa(int(p.Spec.ProjectId)),
	}

	serviceAccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Name,
			Namespace: p.Namespace,
			Labels:    labels,
		},
	}
	// Set Memcached instance as the owner and controller
	ctrl.SetControllerReference(p, serviceAccount, r.Scheme)
	return serviceAccount
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&launchboxiov1alpha1.Project{}).
		Owns(&v1.ServiceAccount{}).
		Complete(r)
}
