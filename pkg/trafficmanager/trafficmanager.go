package trafficmanager

import (
	launchboxiov1alpha1 "github.com/launchboxio/launchbox/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

type ReconcileStatus struct {
	Result    ctrl.Result
	Revisions []launchboxiov1alpha1.ActiveRevisionStatus
}

type TrafficManagerConfig struct {
	Interval time.Duration
}

type TrafficManager struct {
	Type   string
	Config TrafficManagerConfig
}

func (tm *TrafficManager) projectDueForEvaluation() bool {
	return true
}

func (tm *TrafficManager) Evaluate(proj *launchboxiov1alpha1.Project, revisions []*launchboxiov1alpha1.Revision) (*ReconcileStatus, error) {
	switch tm.Type {
	case "canary":
		return tm.canary(proj, revisions)
	case "blue_green":
		return tm.blueGreen(proj, revisions)
	case "in_place":
		return tm.inPlace(proj, revisions)
	}
	return
}
