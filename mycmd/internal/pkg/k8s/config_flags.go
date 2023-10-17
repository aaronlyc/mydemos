package k8s

import (
	"github.com/spf13/pflag"
	utilpointer "k8s.io/utils/pointer"
)

const (
	flagNamespace = "namespace"
)

// ConfigFlags composes the set of values necessary
// for obtaining a REST client config
type ConfigFlags struct {
	KubeConfig *string
	Namespace  *string
}

// AddFlags binds client configuration flags to a given flagset
func (f *ConfigFlags) AddFlags(flags *pflag.FlagSet) {
	if f.KubeConfig != nil {
		flags.StringVar(f.KubeConfig, "kubeconfig", *f.KubeConfig, "Path to the kubeconfig file to use for CLI requests.")
	}
	if f.Namespace != nil {
		flags.StringVarP(f.Namespace, flagNamespace, "n", *f.Namespace, "If present, the namespace scope for this CLI request")
	}
}

// NewConfigFlags returns ConfigFlags with default values set
func NewConfigFlags() *ConfigFlags {
	return &ConfigFlags{
		KubeConfig: utilpointer.String(""),

		Namespace: utilpointer.String(""),
	}
}
