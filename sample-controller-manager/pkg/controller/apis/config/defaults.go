package config

import (
	sampleconfig "mydemos/sample-controller-manager/pkg/controller/sample/config"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cmconfig "k8s.io/controller-manager/config"
)

// 设置默认值
func SetDefault_ControllerManagerConfiguration(obj *ControllerManagerConfiguration) {
	// set default value for generic
	RecommendedDefaultGenericControllerManagerConfiguration(&obj.Generic)

	// TODO: 这里设置用户自定义的默认值
	sampleconfig.RecommendedDefaultSampleControllerConfiguration(&obj.Samplecontroller)
}

func RecommendedDefaultGenericControllerManagerConfiguration(obj *cmconfig.GenericControllerManagerConfiguration) {
	// set default value for client connection
	if obj.ClientConnection.QPS == 0.0 {
		obj.ClientConnection.QPS = 20.0
	}
	if obj.ClientConnection.Burst == 0 {
		obj.ClientConnection.Burst = 30
	}
	if len(obj.ClientConnection.ContentType) == 0 {
		obj.ClientConnection.ContentType = "application/vnd.kubernetes.protobuf"
	}

	zero := metav1.Duration{}
	if obj.Address == "" {
		obj.Address = "0.0.0.0"
	}
	if obj.MinResyncPeriod == zero {
		obj.MinResyncPeriod = metav1.Duration{Duration: 12 * time.Hour}
	}
	if obj.ControllerStartInterval == zero {
		obj.ControllerStartInterval = metav1.Duration{Duration: 0 * time.Second}
	}
	if len(obj.Controllers) == 0 {
		obj.Controllers = []string{"*"}
	}

	// set the default LeaderElectionConfiguration options
	if len(obj.LeaderElection.ResourceLock) == 0 {
		// Use lease-based leader election to reduce cost.
		// We migrated for EndpointsLease lock in 1.17 and starting in 1.20 we
		// migrated to Lease lock.
		obj.LeaderElection.ResourceLock = "leases"
	}
	if obj.LeaderElection.LeaseDuration == zero {
		obj.LeaderElection.LeaseDuration = metav1.Duration{Duration: 15 * time.Second}
	}
	if obj.LeaderElection.RenewDeadline == zero {
		obj.LeaderElection.RenewDeadline = metav1.Duration{Duration: 10 * time.Second}
	}
	if obj.LeaderElection.RetryPeriod == zero {
		obj.LeaderElection.RetryPeriod = metav1.Duration{Duration: 2 * time.Second}
	}

	if obj.LeaderElection.ResourceName == "" {
		obj.LeaderElection.ResourceName = "sample-controller-manager"
	}
	if obj.LeaderElection.ResourceNamespace == "" {
		obj.LeaderElection.ResourceNamespace = "sample-system"
	}
}
