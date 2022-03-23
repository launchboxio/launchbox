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
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strconv"

	launchboxiov1alpha1 "github.com/launchboxio/launchbox/api/v1alpha1"
)

var (
	protocolMap = map[string]v1.Protocol{
		"tcp":  v1.ProtocolTCP,
		"udp":  v1.ProtocolUDP,
		"sctp": v1.ProtocolSCTP,
	}
)

// RevisionReconciler reconciles a Revision object
type RevisionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=launchbox.io,resources=revisions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=launchbox.io,resources=revisions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=launchbox.io,resources=revisions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Revision object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *RevisionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	out := log.FromContext(ctx)

	revision := &launchboxiov1alpha1.Revision{}
	err := r.Get(ctx, req.NamespacedName, revision)
	if err != nil {
		if errors.IsNotFound(err) {
			out.Info("Revision resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		out.Error(err, "Failed to get Revision")
		return ctrl.Result{}, err
	}

	project := &launchboxiov1alpha1.Project{}
	err = r.Get(ctx, types.NamespacedName{Name: revision.Spec.ProjectName, Namespace: revision.Namespace}, project)
	if err != nil {
		if errors.IsNotFound(err) {
			out.Info("Project resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		out.Error(err, "Failed to get Project")
		return ctrl.Result{}, err
	}

	// Create the service
	foundService := &v1.Service{}
	err = r.Get(ctx, types.NamespacedName{Name: revision.Name, Namespace: revision.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		service := r.serviceForRevision(revision, project)
		out.Info("Creating a new Service", "Service.Namespace", revision.Namespace, "Service.Name", revision.Name)
		err = r.Create(ctx, service)
		if err != nil {
			out.Error(err, "Failed to create new service", "Service.Namespace", revision.Namespace, "Service.Name", revision.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		out.Error(err, "Failed to get service")
		return ctrl.Result{}, err
	}

	if foundService.Name != revision.Status.Service {
		fmt.Println("Updating revision with new service information")
		revision.Status.Service = foundService.Name
		err := r.Status().Update(ctx, revision)
		if err != nil {
			out.Error(err, "Failed to update Revision status")
			return ctrl.Result{}, err
		}
	}

	// TODO: Create the deployment
	foundDeployment := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: revision.Name, Namespace: revision.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		deployment := r.deploymentForRevision(revision, project)
		out.Info("Creating a new Deployment", "Deployment.Namespace", revision.Namespace, "Deployment.Name", revision.Name)
		err = r.Create(ctx, deployment)
		if err != nil {
			out.Error(err, "Failed to create new deployment", "Deployment.Namespace", revision.Namespace, "Deployment.Name", revision.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		out.Error(err, "Failed to get deployment")
		return ctrl.Result{}, err
	}

	if foundDeployment.Name != revision.Status.Deployment {
		fmt.Println("Updating revision with new deployment information")
		revision.Status.Deployment = foundDeployment.Name
		err := r.Status().Update(ctx, revision)
		if err != nil {
			out.Error(err, "Failed to update deployment status")
			return ctrl.Result{}, err
		}
	}

	// TODO: If autoscaling.enabled, create the HorizontalPodAutoscaler

	// TODO: Create required OSM metrics resources

	// TODO: Transition traffic
	// if first revision
	// 		set traffic to 100
	//		update all statuses appropriately
	// else
	// 	 	ensure traffic 100 on v-1
	// 		for each step
	// 			set new traffic amount / percentage
	//			wait for specified time
	//			monitor new service and compare metrics
	// 			if success
	//				update statuses
	//              continue
	//			else
	//				revert traffic to 100 on v-1
	//				set statuses to 'rolled back'
	//

	return ctrl.Result{}, nil
}

func (r *RevisionReconciler) serviceForRevision(rev *launchboxiov1alpha1.Revision, proj *launchboxiov1alpha1.Project) *v1.Service {
	labels := map[string]string{
		"launchbox.io/project.id":  rev.Spec.ProjectName,
		"launchbox.io/revision.id": rev.Name,
	}
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rev.Name,
			Namespace: rev.Namespace,
			Labels:    labels,
			//Annotations: annotations,
		},
		Spec: v1.ServiceSpec{
			Ports: crdPortsToServicePorts(rev.Spec.Ports),
			Selector: map[string]string{
				"launchbox.io/project.id":  strconv.Itoa(int(proj.Spec.ProjectId)),
				"launchbox.io/revision.id": rev.Name,
			},
		},
	}
	ctrl.SetControllerReference(rev, service, r.Scheme)
	return service
}

func (r *RevisionReconciler) deploymentForRevision(rev *launchboxiov1alpha1.Revision, proj *launchboxiov1alpha1.Project) *appsv1.Deployment {
	labels := map[string]string{
		"launchbox.io/project.id":  rev.Spec.ProjectName,
		"launchbox.io/revision.id": rev.Name,
	}
	ports := []v1.ContainerPort{}
	for _, port := range rev.Spec.Ports {
		containerPort := v1.ContainerPort{
			Name:          port.Name,
			ContainerPort: port.Port,
			//Protocol:      protocolMap[port.Protocol],
		}
		ports = append(ports, containerPort)
	}
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rev.Name,
			Namespace: rev.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: nil,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"launchbox.io/revision.id": rev.Name,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"launchbox.io/revision.id": rev.Name,
						"launchbox.io/project.id":  strconv.Itoa(int(proj.Spec.ProjectId)),
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Name:  "app",
						Image: rev.Spec.Image,
						Ports: ports,
						// TODO: Obviously we're missing a ton here...
					}},
					ServiceAccountName: proj.Status.ServiceAccount,
				},
			},
			Strategy:                appsv1.DeploymentStrategy{},
			MinReadySeconds:         0,
			RevisionHistoryLimit:    nil,
			Paused:                  false,
			ProgressDeadlineSeconds: nil,
		},
	}
	ctrl.SetControllerReference(rev, dep, r.Scheme)
	return dep
}

// SetupWithManager sets up the controller with the Manager.
func (r *RevisionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&launchboxiov1alpha1.Revision{}).
		Owns(&v1.Service{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
