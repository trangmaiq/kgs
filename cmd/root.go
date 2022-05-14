package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/trangmaiq/kgs/cmd/generate"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "kgs",
	}

	cmd.AddCommand(generate.NewGenerateCmd())

	return cmd
}

func Execute() {
	cmd := NewRootCmd()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
