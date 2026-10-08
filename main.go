package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

// version is overridden at build time via -ldflags "-X main.version=...".
var version = "dev"

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	// Standard kubectl config flags (--context, --namespace, --kubeconfig, ...).
	configFlags := genericclioptions.NewConfigFlags(true)
	streams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}

	cmd := &cobra.Command{
		Use:          "kubectl resources-missing [flags]",
		Short:        "Find objects with no resources block (missing cpu/memory requests & limits)",
		Version:      version,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _, err := configFlags.ToRawKubeConfigLoader().Namespace()
			if err != nil {
				return err
			}
			fmt.Fprintf(streams.Out, "resources-missing: not implemented yet (namespace=%q)\n", namespace)
			// TODO: build a client from configFlags.ToRESTConfig(), scan
			// workloads, and report containers missing resource requests/limits.
			return nil
		},
	}

	configFlags.AddFlags(cmd.PersistentFlags())
	return cmd
}
