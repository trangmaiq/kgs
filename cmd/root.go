package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/trangmaiq/kgs/cmd/generate"
	"github.com/trangmaiq/kgs/cmd/serve"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "kgs",
	}

	cmd.AddCommand(generate.NewGenerateCmd())
	cmd.AddCommand(serve.NewServeCmd())

	return cmd
}

func Execute() {
	cmd := NewRootCmd()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
