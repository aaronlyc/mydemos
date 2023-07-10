package options

import (
	"fmt"
	"mydemos/sample-controller-manager/pkg/cluster/ports"
	apisconfig "mydemos/sample-controller-manager/pkg/controller/apis/config"
	"net"

	v1 "k8s.io/api/core/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	apiserveroptions "k8s.io/apiserver/pkg/server/options"

	appconfig "mydemos/sample-controller-manager/cmd/app/config"

	clientgokubescheme "k8s.io/client-go/kubernetes/scheme"

	clientset "k8s.io/client-go/kubernetes"

	utilfeature "k8s.io/apiserver/pkg/util/feature"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
	"k8s.io/component-base/metrics"
	cmoptions "k8s.io/controller-manager/options"
	netutils "k8s.io/utils/net"
)

const (
	// SampleControllerManagerUserAgent is the userAgent name when starting sample controller managers.
	SampleControllerManagerUserAgent = "sample-controller-manager"
)

type ControllerManagerOptions struct {
	Generic *cmoptions.GenericControllerManagerConfigurationOptions

	// TODO: here is the custormer controller options
	SampleController *SampleControllerOptions

	SecureServing *apiserveroptions.SecureServingOptionsWithLoopback
	// Authentication *apiserveroptions.DelegatingAuthenticationOptions
	// Authorization  *apiserveroptions.DelegatingAuthorizationOptions
	Metrics *metrics.Options
	Logs    *logs.Options

	Master                      string
	ShowHiddenMetricsForVersion string
}

func NewControllerManagerOptions() (*ControllerManagerOptions, error) {
	componentConfig, err := NewDefaultComponentConfig()
	if err != nil {
		return nil, err
	}

	s := ControllerManagerOptions{
		Generic:       cmoptions.NewGenericControllerManagerConfigurationOptions(&componentConfig.Generic),
		SecureServing: apiserveroptions.NewSecureServingOptions().WithLoopback(),
		// Authentication: apiserveroptions.NewDelegatingAuthenticationOptions(),
		// Authorization:  apiserveroptions.NewDelegatingAuthorizationOptions(),
		Metrics: metrics.NewOptions(),
		Logs:    logs.NewOptions(),

		SampleController: &SampleControllerOptions{
			&componentConfig.Samplecontroller,
		},
	}

	// s.Authentication.RemoteKubeConfigFileOptional = true
	// s.Authorization.RemoteKubeConfigFileOptional = true

	// Set the PairName but leave certificate directory blank to generate in-memory by default
	s.SecureServing.ServerCert.CertDirectory = ""
	s.SecureServing.ServerCert.PairName = "sample-controller-manager"
	s.SecureServing.BindPort = ports.SampleControllerManagerPort

	return &s, nil
}

// Flags returns flags for a specific Controller by section name
func (o *ControllerManagerOptions) Flags(allControllers []string, disabledByDefaultControllers []string) cliflag.NamedFlagSets {
	fss := cliflag.NamedFlagSets{}
	o.Generic.AddFlags(&fss, allControllers, disabledByDefaultControllers)
	o.Metrics.AddFlags(fss.FlagSet("metrics"))
	logsapi.AddFlags(o.Logs, fss.FlagSet("logs"))

	// TODO: 用户的自定义flags
	o.SampleController.AddFlags(fss.FlagSet("sample-controller"))

	fs := fss.FlagSet("misc")
	fs.StringVar(&o.Master, "master", o.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig).")
	fs.StringVar(&o.Generic.ClientConnection.Kubeconfig, "kubeconfig", o.Generic.ClientConnection.Kubeconfig, "Path to kubeconfig file with authorization and master location information (the master location can be overridden by the master flag).")
	utilfeature.DefaultMutableFeatureGate.AddFlag(fss.FlagSet("generic"))

	return fss
}

func (o *ControllerManagerOptions) Validate(allControllers []string, disabledByDefaultControllers []string) error {
	var errs []error

	errs = append(errs, o.Generic.Validate(allControllers, disabledByDefaultControllers)...)
	errs = append(errs, o.SampleController.Validate()...)

	return utilerrors.NewAggregate(errs)
}

func (o *ControllerManagerOptions) ApplyTo(c *appconfig.Config) error {
	if err := o.Generic.ApplyTo(&c.ComponentConfig.Generic); err != nil {
		return err
	}
	if err := o.SecureServing.ApplyTo(&c.SecureServing, &c.LoopbackClientConfig); err != nil {
		return err
	}
	// if o.SecureServing.BindPort != 0 || o.SecureServing.Listener != nil {
	// 	if err := o.Authentication.ApplyTo(&c.Authentication, c.SecureServing, nil); err != nil {
	// 		return err
	// 	}
	// 	if err := o.Authorization.ApplyTo(&c.Authorization); err != nil {
	// 		return err
	// 	}
	// }

	// TODO: custom controller applyto
	if err := o.SampleController.ApplyTo(&c.ComponentConfig.Samplecontroller); err != nil {
		return err
	}

	return nil
}

// NewDefaultComponentConfig returns kube-controller manager configuration object.
func NewDefaultComponentConfig() (apisconfig.ControllerManagerConfiguration, error) {
	result := apisconfig.ControllerManagerConfiguration{}
	apisconfig.SetDefault_ControllerManagerConfiguration(&result)
	return result, nil
}

// Config return a controller manager config objective
func (o ControllerManagerOptions) Config(allControllers []string, disabledByDefaultControllers []string) (*appconfig.Config, error) {
	if err := o.Validate(allControllers, disabledByDefaultControllers); err != nil {
		return nil, err
	}

	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{netutils.ParseIPSloppy("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	kubeconfig, err := clientcmd.BuildConfigFromFlags(o.Master, o.Generic.ClientConnection.Kubeconfig)
	if err != nil {
		return nil, err
	}
	kubeconfig.DisableCompression = true
	kubeconfig.ContentConfig.AcceptContentTypes = o.Generic.ClientConnection.AcceptContentTypes
	kubeconfig.ContentConfig.ContentType = o.Generic.ClientConnection.ContentType
	kubeconfig.QPS = o.Generic.ClientConnection.QPS
	kubeconfig.Burst = int(o.Generic.ClientConnection.Burst)

	client, err := clientset.NewForConfig(restclient.AddUserAgent(kubeconfig, SampleControllerManagerUserAgent))
	if err != nil {
		return nil, err
	}

	eventBroadcaster := record.NewBroadcaster()
	eventRecorder := eventBroadcaster.NewRecorder(clientgokubescheme.Scheme, v1.EventSource{Component: SampleControllerManagerUserAgent})

	c := &appconfig.Config{
		Client:           client,
		Kubeconfig:       kubeconfig,
		EventBroadcaster: eventBroadcaster,
		EventRecorder:    eventRecorder,
	}
	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}
	o.Metrics.Apply()

	return c, nil
}
