package trafficmanager

import (
	"errors"
	launchboxiov1alpha1 "github.com/launchboxio/launchbox/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Canary checks the requested revisions, analyzes the traffic and statuses
// in the cluster, and returns a result set of desired trafficPercentags
func (tm *TrafficManager) canary(proj *launchboxiov1alpha1.Project, revisions []*launchboxiov1alpha1.Revision) (*ReconcileStatus, error) {
	status := &ReconcileStatus{
		Result:    ctrl.Result{},
		Revisions: []launchboxiov1alpha1.ActiveRevisionStatus{},
	}

	if len(revisions) == 0 {
		return status, nil
	}

	if len(revisions) > 2 {
		return status, errors.New("Only 2 revisions at a time are supported")
	}

	// If the project is not ready to have its revisions evaluated,
	// we simply short circuit, and try again later
	if !tm.projectDueForEvaluation() {
		// Requeue the project to check again momentarily
		// TODO: Should be a delta between last check and interval, not interval right out
		status.Result = ctrl.Result{RequeueAfter: tm.Config.Interval}
		return status, nil
	}

	// If we have only one revision, we want to set it to 100%, and return
	// a slice with that single revision
	if len(revisions) == 1 {
		status.Revisions = []launchboxiov1alpha1.ActiveRevisionStatus{{
			RevisionId:        revisions[0].Name,
			Status:            "primary",
			TrafficPercentage: 100,
		}}
		return status, nil
	}

	// The meat and potatoes. We have N revisions. We need to run through
	// them, (eventually) evaluate metrics / performance, and increment
	// the traffic percentage appropriately. We also need to support rollback
	// in the event that something is failing. At the moment, N cannot be
	// > 2, we only support canary-ing between 2 releases

	return status, nil
}
