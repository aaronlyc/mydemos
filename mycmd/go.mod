module mydemos/mycmd

go 1.16

require (
	github.com/lithammer/dedent v1.1.0
	github.com/moby/term v0.5.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.27.3
	k8s.io/client-go v0.0.0
	k8s.io/klog/v2 v2.120.1
	sigs.k8s.io/yaml v1.4.0
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.27.3
	k8s.io/client-go => k8s.io/client-go v0.27.3
)
