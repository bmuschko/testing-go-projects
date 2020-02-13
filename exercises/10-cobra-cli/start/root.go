package main

import (
	"github.com/spf13/cobra"
)

func NewRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "calc",
		Short: "A calculator program using Cobra.",
	}

	flags := cmd.PersistentFlags()
	_ = flags.Parse(args)
	out := cmd.OutOrStdout()

	cmd.AddCommand(
		newAddCmd(out),
	)

	return cmd
}
