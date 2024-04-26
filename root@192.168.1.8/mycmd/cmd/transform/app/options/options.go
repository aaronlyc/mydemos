package options

import (
	"mydemos/mycmd/pkg/cli"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Options struct {
	// The default values.
	Kubeconfig string
	// ConfigFile is the location of the scheduler server's configuration file.
	ConfigFile string

	// WriteConfigTo is the path where the default configuration will be written.
	WriteConfigTo string

	Filenames []string

	// Logs  *logs.Options
	Flags *cli.NamedFlagSets
}

func NewOptions() *Options {
	o := &Options{
		Kubeconfig:    "~/.kube/config",
		ConfigFile:    "config.yaml",
		WriteConfigTo: "./config/",
		// Logs:          logs.NewOptions(),
	}

	o.initFlags("transform")

	return o
}

func (o *Options) initFlags(name string) {
	nfs := cli.NamedFlagSets{}
	fs := nfs.FlagSet(name)

	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "Path to kubeconfig file with authorization and master location information (the master location can be overridden by the master flag).")
	fs.StringVar(&o.ConfigFile, "config", o.ConfigFile, "The path to the configuration file.")
	fs.StringVar(&o.WriteConfigTo, "write-config-to", o.WriteConfigTo, "If set, write the configuration values to this file and exit.")
	AddJsonFilenameFlag(fs, &o.Filenames, "Filename, directory, or URL to files that contains the last-applied-configuration annotations")

	o.Flags = &nfs
}

func AddJsonFilenameFlag(flags *pflag.FlagSet, value *[]string, usage string) {
	flags.StringSliceVarP(value, "filename", "f", *value, usage)
	annotations := make([]string, 0, len(fileExtensions))
	for _, ext := range fileExtensions {
		annotations = append(annotations, strings.TrimLeft(ext, "."))
	}
	flags.SetAnnotation("filename", cobra.BashCompFilenameExt, annotations)
}

// var fileExtensions = []string{".json", ".yaml", ".yml"}
var fileExtensions = []string{".yaml", ".yml"}
