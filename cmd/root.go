package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unnecessary-logic/resources-missing-plugin/internal/output"
	"github.com/unnecessary-logic/resources-missing-plugin/internal/scan"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

var version = "dev"

func RootCmd() *cobra.Command {
	configFlags := genericclioptions.NewConfigFlags(true)

	cmd := &cobra.Command{
		Use:          "kubectl resources-missing [flags]",
		Short:        "Find objects with no resources block (missing cpu/memory requests & limits)",
		Version:      version,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace := ""
			if configFlags.Namespace != nil {
				namespace = *configFlags.Namespace
			}

			config, err := configFlags.ToRESTConfig()
			if err != nil {
				return fmt.Errorf("error loading kubeconfig: %s", err)
			}

			client, err := kubernetes.NewForConfig(config)
			if err != nil {
				return fmt.Errorf("error creating client: %s", err)
			}

			out := cmd.OutOrStdout()
			if namespace == "" {
				output.ScanningAll(out)
			}

			results, err := scan.RunScan(cmd.Context(), client, namespace)
			if err != nil {
				return fmt.Errorf("error listing pods: %s", err)
			}
			output.Discovery(out, results)

			return nil
		},
	}

	configFlags.AddFlags(cmd.PersistentFlags())
	return cmd
}
