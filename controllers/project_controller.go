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
	"github.com/go-logr/logr"
	launchboxiov1alpha1 "github.com/launchboxio/launchbox/api/v1alpha1"
	osmv1 "github.com/openservicemesh/osm/pkg/apis/policy/v1alpha1"
	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha4"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"strconv"
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

	out.Info("Reconcile loop for Project", "Project.Namespace", req.Namespace, "Project.Name", req.Name)
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

	// Create the serviceAccount
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

	if found.Name != project.Status.ServiceAccount {
		fmt.Println("Updating project with new service account information")
		project.Status.ServiceAccount = found.Name
		err := r.Status().Update(ctx, project)
		if err != nil {
			out.Error(err, "Failed to update Project status")
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// Create the root service that manages traffic
	foundService := &v1.Service{}
	err = r.Get(ctx, types.NamespacedName{Name: project.Name, Namespace: project.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		service := r.serviceForProject(project)
		out.Info("Creating a new Service", "Service.Namespace", project.Namespace, "Service.Name", project.Name)
		err = r.Create(ctx, service)
		if err != nil {
			out.Error(err, "Failed to create new service", "Service.Namespace", project.Namespace, "Service.Name", project.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		out.Error(err, "Failed to get service")
		return ctrl.Result{}, err
	}

	if foundService.Name != project.Status.RootService {
		fmt.Println("Updating project with new service information")
		project.Status.RootService = foundService.Name
		err := r.Status().Update(ctx, project)
		if err != nil {
			out.Error(err, "Failed to update Project status")
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}

	foundIngress := &v12.Ingress{}
	err = r.Get(ctx, types.NamespacedName{Name: project.Name, Namespace: project.Namespace}, foundIngress)
	if err != nil && errors.IsNotFound(err) {
		ingress := r.ingressForProject(project)
		out.Info("Creating a new Ingress", "Ingress.Namespace", project.Namespace, "Ingress.Name", project.Name)
		err = r.Create(ctx, ingress)
		if err != nil {
			out.Error(err, "Failed to create new ingress", "Ingress.Namespace", project.Namespace, "Ingress.Name", project.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		out.Error(err, "Failed to get ingress")
		return ctrl.Result{}, err
	}

	foundBackend := &osmv1.IngressBackend{}
	err = r.Get(ctx, types.NamespacedName{Name: project.Name, Namespace: project.Namespace}, foundBackend)
	if err != nil && errors.IsNotFound(err) {
		ingress := r.ingressBackendForProject(project)
		out.Info("Creating a new IngressBackend", "IngressBackend.Namespace", project.Namespace, "IngressBackend.Name", project.Name)
		err = r.Create(ctx, ingress)
		if err != nil {
			out.Error(err, "Failed to create new ingress backend", "IngressBackend.Namespace", project.Namespace, "IngressBackend.Name", project.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		out.Error(err, "Failed to get ingress backend")
		return ctrl.Result{}, err
	}

	out.Info("Analyzing revisions", "Count", len(project.Status.ActiveRevisions))

	// TODO: We should lock the project resource, so we don't jack anything up?
	wasUpdated, res, err := r.manageRevisionTrafficSplits(project, out)
	// TODO: Unlock
	if err != nil {
		out.Error(err, "Failed allocating revision traffic")
		return res, err
	}
	if wasUpdated {
		return ctrl.Result{Requeue: true}, nil
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&launchboxiov1alpha1.Project{}).
		Owns(&v1.ServiceAccount{}).
		Owns(&v1.Service{}).
		Owns(&v12.Ingress{}).
		Owns(&osmv1.IngressBackend{}).
		Complete(r)
}

func (p *ProjectReconciler) manageRevisionTrafficSplits(proj *launchboxiov1alpha1.Project, out logr.Logger) (bool, reconcile.Result, error) {
	// Easy first case. We have one live revision. Short circuit, set to 100%, and return
	revisions := proj.Status.ActiveRevisions
	if len(revisions) == 1 {
		out.Info("Single active revision found, setting to 100% traffic")
		revision := revisions[0]
		revision.TrafficPercentage = 100
		revision.Status = "primary"
		return p.ensureTrafficSplit(context.TODO(), proj, []launchboxiov1alpha1.ActiveRevisionStatus{revision})
	}
	return false, ctrl.Result{}, nil
}

func (p *ProjectReconciler) ensureTrafficSplit(ctx context.Context, proj *launchboxiov1alpha1.Project, revs []launchboxiov1alpha1.ActiveRevisionStatus) (bool, reconcile.Result, error) {
	split := &v1alpha4.TrafficSplit{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proj.Name,
			Namespace: proj.Namespace,
		},
		Spec: v1alpha4.TrafficSplitSpec{
			Service:  fmt.Sprintf("%s.%s.svc.cluster.local", proj.Name, proj.Namespace),
			Backends: []v1alpha4.TrafficSplitBackend{},
		},
	}
	for _, rev := range revs {
		split.Spec.Backends = append(split.Spec.Backends, v1alpha4.TrafficSplitBackend{
			Service: rev.RevisionId,
			Weight:  int(rev.TrafficPercentage),
		})
	}

	// Upsert the split
	foundSplit := &v1alpha4.TrafficSplit{}
	err := p.Get(ctx, types.NamespacedName{Name: proj.Name, Namespace: proj.Namespace}, foundSplit)
	if err != nil && errors.IsNotFound(err) {
		err = p.Create(ctx, split)
		if err != nil {
			return true, ctrl.Result{}, err
		}
		return true, ctrl.Result{}, nil
	} else if err != nil {
		return true, ctrl.Result{}, err
	}

	// Ensure split is updated
	if !reflect.DeepEqual(foundSplit.Spec.Backends, split.Spec.Backends) {
		err := p.Status().Update(ctx, split)
		if err != nil {
			return true, ctrl.Result{}, err
		}
		return true, ctrl.Result{Requeue: true}, nil
	}

	// Ensure project is updated
	if !reflect.DeepEqual(proj.Status.ActiveRevisions, revs) {
		proj.Status.ActiveRevisions = revs
		err := p.Status().Update(ctx, proj)
		if err != nil {
			return true, ctrl.Result{}, err
		}
		return true, ctrl.Result{Requeue: true}, nil
	}

	return false, ctrl.Result{}, nil
}

func (p *ProjectReconciler) serviceAccountForProject(proj *launchboxiov1alpha1.Project) *v1.ServiceAccount {
	labels := map[string]string{
		"launchbox.io/application.id": strconv.Itoa(int(proj.Spec.ApplicationId)),
		"launchbox.io/project.id":     strconv.Itoa(int(proj.Spec.ProjectId)),
	}

	serviceAccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proj.Name,
			Namespace: proj.Namespace,
			Labels:    labels,
		},
	}
	ctrl.SetControllerReference(proj, serviceAccount, p.Scheme)
	return serviceAccount
}

func (r *ProjectReconciler) serviceForProject(proj *launchboxiov1alpha1.Project) *v1.Service {
	labels := map[string]string{
		"launchbox.io/project.id":     proj.Name,
		"launchbox.io/application.id": strconv.Itoa(int(proj.Spec.ApplicationId)),
	}
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proj.Name,
			Namespace: proj.Namespace,
			Labels:    labels,
		},
		Spec: v1.ServiceSpec{
			Ports: crdPortsToServicePorts(proj.Spec.Ports),
			Selector: map[string]string{
				"launchbox.io/project.id": strconv.Itoa(int(proj.Spec.ProjectId)),
			},
		},
	}
	ctrl.SetControllerReference(proj, service, r.Scheme)
	return service
}

func (r *ProjectReconciler) ingressForProject(proj *launchboxiov1alpha1.Project) *v12.Ingress {
	ingressClassName := "nginx"
	pathPrefix := v12.PathTypePrefix
	ingress := &v12.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proj.Name,
			Namespace: proj.Namespace,
		},
		Spec: v12.IngressSpec{
			IngressClassName: &ingressClassName,
			Rules: []v12.IngressRule{{
				// TODO: Generate more friendly subdomains, instead of using the namespace
				Host: fmt.Sprintf("%s.launchbox.local", proj.Namespace),
				IngressRuleValue: v12.IngressRuleValue{
					HTTP: &v12.HTTPIngressRuleValue{
						Paths: []v12.HTTPIngressPath{{
							Path:     "/",
							PathType: &pathPrefix,
							Backend: v12.IngressBackend{
								Service: &v12.IngressServiceBackend{
									Name: proj.Name,
									Port: v12.ServiceBackendPort{
										Number: proj.Spec.Ports[0].Port,
									},
								},
							},
						}},
					},
				},
			}},
		},
	}
	ctrl.SetControllerReference(proj, ingress, r.Scheme)
	return ingress
}

func (r *ProjectReconciler) ingressBackendForProject(proj *launchboxiov1alpha1.Project) *osmv1.IngressBackend {
	backend := &osmv1.IngressBackend{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proj.Name,
			Namespace: proj.Namespace,
		},
		Spec: osmv1.IngressBackendSpec{
			Backends: []osmv1.BackendSpec{{
				Name: proj.Name,
				Port: osmv1.PortSpec{
					Number:   int(proj.Spec.Ports[0].Port),
					Protocol: "http",
				},
			}},
			Sources: []osmv1.IngressSourceSpec{{
				Kind:      "Service",
				Namespace: "ingress-nginx",
				Name:      "ingress-nginx-controller",
			}},
		},
	}
	ctrl.SetControllerReference(proj, backend, r.Scheme)
	return backend
}

func crdPortsToServicePorts(ports []launchboxiov1alpha1.Port) []v1.ServicePort {
	res := []v1.ServicePort{}
	for _, port := range ports {
		item := v1.ServicePort{
			Name:     port.Name,
			Port:     port.Port,
			NodePort: 0,
		}
		if port.TargetPort != "" {
			item.TargetPort = intstr.IntOrString{
				Type:   intstr.String,
				StrVal: port.TargetPort,
			}
		}
		res = append(res, item)
	}
	return res
}
