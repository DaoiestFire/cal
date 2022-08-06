package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GenerateCalCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "cal [subCommand] arg1 arg2 [options]",
		Long: `cal is simple calculator app based on cobra library
it supports add/sub/div/mul four operation.
hurry to try it!`,
		Version: "0.1",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return fmt.Errorf("cal cmd only accepts flag \"-v\" or \"--version\"")
			}
			return nil
		},
	}

	return rootCmd
}
