package cmd

import (
	"mydemos/mycmd/internal/pkg/k8s"

	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/cmd/plugin"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

type MyUpgradeOptions struct {
	// PluginHandler PluginHandler
	// Arguments   []string
	ConfigFlags *k8s.ConfigFlags

	k8s.IOStreams
}

func NewDefaultMyUpgradeCommand() *cobra.Command {

}

func NewDefaultMyUpgradeCommandWithArgs(o MyUpgradeOptions) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "myupgrade",
		Short: i18n.T("myupgrade controls the appnest cluster manager"),
		Long: templates.LongDesc(`
		myupgrade controls the appnest cluster manager.`),
		Run: runHelp,
		// Hook before and after Run initialize and write profiles to disk,
		// respectively.
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// rest.SetDefaultWarningHandler(warningHandler)

			if cmd.Name() == cobra.ShellCompRequestCmd {
				// This is the __complete or __completeNoDesc command which
				// indicates shell completion has been requested.
				plugin.SetupPluginCompletion(cmd, args)
			}

			return nil
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			// if err := flushProfiling(); err != nil {
			// 	return err
			// }
			// if warningsAsErrors {
			// 	count := warningHandler.WarningCount()
			// 	switch count {
			// 	case 0:
			// 		// no warnings
			// 	case 1:
			// 		return fmt.Errorf("%d warning received", count)
			// 	default:
			// 		return fmt.Errorf("%d warnings received", count)
			// 	}
			// }
			return nil
		},
	}
	// From this point and forward we get warnings on flags that contain "_" separators
	// when adding them with hyphen instead of the original name.
	// cmds.SetGlobalNormalizationFunc(cliflag.WarnWordSepNormalizeFunc)

	flags := cmds.PersistentFlags()
	kubeConfigFlags := o.ConfigFlags
	if kubeConfigFlags == nil {
		kubeConfigFlags = k8s.NewConfigFlags()
	}
	kubeConfigFlags.AddFlags(flags)

}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
