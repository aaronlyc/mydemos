/*
First, we have some *package-level* markers that denote that there are
Kubernetes objects in this package, and that this package represents the group
`batch.tutorial.kubebuilder.io`. The `object` generator makes use of the
former, while the latter is used by the CRD generator to generate the right
metadata for the CRDs it creates from this package.
*/

// Package v1 contains API Schema definitions for the batch v1 API group
// +kubebuilder:object:generate=true
// +groupName=batch.tutorial.kubebuilder.io
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

/*
Then, we have the commonly useful variables that help us set up our Scheme.
Since we need to use all the types in this package in our controller, it's
helpful (and the convention) to have a convenient method to add all the types to
some other `Scheme`. SchemeBuilder makes this easy for us.
*/
var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "batch.tutorial.kubebuilder.io", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
