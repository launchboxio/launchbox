package trafficmanager

import (
	"errors"
	launchboxiov1alpha1 "github.com/launchboxio/launchbox/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (tm *TrafficManager) inPlace(proj *launchboxiov1alpha1.Project, revisions []*launchboxiov1alpha1.Revision) (*ReconcileStatus, error) {
	status := &ReconcileStatus{
		Result:    ctrl.Result{},
		Revisions: []launchboxiov1alpha1.ActiveRevisionStatus{},
	}

	return status, errors.New("InPlace deployments not yet supported")
}
