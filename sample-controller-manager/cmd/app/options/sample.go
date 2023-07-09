package options

import (
	sampleconfig "mydemos/sample-controller-manager/pkg/controller/sample/config"

	"github.com/spf13/pflag"
)

type SampleControllerOptions struct {
	*sampleconfig.SamplecontrollerConfiguration
}

// AddFlags adds flags related to SampleController for controller manager to the specified FlagSet.
func (o *SampleControllerOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.StringVar(&o.ControllerName, "sample-controller-name", o.ControllerName, "The name for sample controller")
	fs.Int32Var(&o.ConcurrentSampleSyncs, "sample-controller-syncs", o.ConcurrentSampleSyncs, "The number of sample objects that are allowed to sync concurrently. Larger number = more responsive samples, but more CPU (and network) load")
}

// ApplyTo fills up DeploymentController config with options.
func (o *SampleControllerOptions) ApplyTo(cfg *sampleconfig.SamplecontrollerConfiguration) error {
	if o == nil {
		return nil
	}

	cfg.ConcurrentSampleSyncs = o.ConcurrentSampleSyncs
	cfg.ControllerName = "sample-controller"

	return nil
}

// Validate checks validation of DeploymentControllerOptions.
func (o *SampleControllerOptions) Validate() []error {
	if o == nil {
		return nil
	}

	errs := []error{}
	return errs
}
