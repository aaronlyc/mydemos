package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *EtcdCluster) AsOwner() metav1.OwnerReference {
	trueVar := true
	return metav1.OwnerReference{
		APIVersion: GroupVersion.String(),
		Kind:       "EtcdCluster",
		Name:       c.Name,
		UID:        c.UID,
		Controller: &trueVar,
	}
}
