package app

import (
	"context"
	"errors"
	"fmt"
	"mydemos/mycmd/cmd/transform/app/options"
	"mydemos/mycmd/pkg/cli"
	"mydemos/mycmd/pkg/signals"
	"mydemos/mycmd/pkg/term"
	"strings"

	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	applyExample = dedent.Dedent(`
		# Apply the configuration in pod.yaml to a pod
		transform -f ./pod.yaml

		# Apply resources from a directory - e.g. dir/*.yaml
		transform -f dir/

		# Apply the YAML passed into stdin to a pod
		cat pod.yaml | transform -f -

		# Apply the configuration from all files that end with '.yaml' - i.e. expand wildcard characters in file names
		transform -f '*.yaml'`)
)

func NewTransformCommand() *cobra.Command {
	opts := options.NewOptions()

	cmd := &cobra.Command{
		Use:   "transform -f FILENAME",
		Short: `Transform files to the supported resource API versions in Kubernetes`,
		Long: dedent.Dedent(`The transform command analyzes the resource configuration manifest file 
specified by <input-file>, and based on the Kubernetes environment's supported resource API versions, 
it transforms the file to ensure compatibility. The transformed manifest is then written to <output-file>.`),
		SilenceErrors:         true,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
		Example:               applyExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCommand(cmd, opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cli.SetUsageAndHelpFunc(cmd, *opts.Flags, cols)

	return cmd
}

// runCommand runs the scheduler.
func runCommand(cmd *cobra.Command, opts *options.Options) error {

	// Activate logging as soon as possible, after that
	// show flags with the final logging configuration.
	// if err := logsapi.ValidateAndApply(opts.Logs, nil); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%v\n", err)
	// 	os.Exit(1)
	// }
	cli.PrintFlags(cmd.Flags())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		stopCh := signals.SetupSignalHandler()
		<-stopCh.Done()
		cancel()
	}()

	return Run(ctx, opts)
}

func Run(ctx context.Context, opts *options.Options) error {

	client, err := createK8sClient(ctx, opts)
	if err != nil {
		return err
	}

	_, apiResourceList, err := client.ServerGroupsAndResources()
	if err != nil {
		return err
	}

	// 输出所有资源信息
	for _, list := range apiResourceList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			return err
		}

		for _, resource := range list.APIResources {
			var builder strings.Builder
			if resource.Name == "" {
				return errors.New("resource name is empty")
			}
			builder.WriteString(resource.Name)

			if gv.Group != "" {
				builder.WriteString(".")
				builder.WriteString(gv.Group)
			}

			gk := builder.String()

			if gv.Version != "" {
				builder.WriteString("/")
				builder.WriteString(gv.Version)
			}
			fmt.Printf(`
apiVersion: %s
gk: %s 
version: %s
			
			`, builder.String(), gk, gv.Version)
		}
	}

	return nil
}

func createK8sClient(_ context.Context, opts *options.Options) (*discovery.DiscoveryClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", opts.Kubeconfig)
	if err != nil {
		return nil, err
	}

	// discovery.NewDiscoveryClientForConfigg函数通过config实例化discoveryClient对象
	return discovery.NewDiscoveryClientForConfig(config)
}
