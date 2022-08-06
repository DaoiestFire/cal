package cmd

import (
	"fmt"
	"ljw/cal/cmd/options"

	"github.com/spf13/cobra"
)

func GenerateCalCmd() *cobra.Command {
	o := &options.Options{}

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

	rootCmd.PersistentFlags().IntVar(&o.ArgOne, "argone", 0, "assign a value to arg1")
	rootCmd.PersistentFlags().IntVar(&o.ArgTwo, "argtwo", 0, "assign a value to arg2")

	rootCmd.AddCommand(GenerateAddSubCommand(o))
	rootCmd.AddCommand(GenerateSubSubCommand(o))
	rootCmd.AddCommand(GenerateMulSubCommand(o))
	rootCmd.AddCommand(GenerateDivSubCommand(o))
	return rootCmd
}

func GenerateAddSubCommand(o *options.Options) *cobra.Command {
	addCmd := &cobra.Command{
		Use: "add",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("result: %d", o.ArgOne+o.ArgTwo)
			return nil
		},
	}

	return addCmd
}

func GenerateSubSubCommand(o *options.Options) *cobra.Command {
	subCmd := &cobra.Command{
		Use: "sub",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("result: %d", o.ArgOne-o.ArgTwo)
			return nil
		},
	}

	return subCmd
}

func GenerateMulSubCommand(o *options.Options) *cobra.Command {
	mulCmd := &cobra.Command{
		Use: "mul",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("result: %d", o.ArgOne*o.ArgTwo)
			return nil
		},
	}

	return mulCmd
}

func GenerateDivSubCommand(o *options.Options) *cobra.Command {
	var precision int
	divCmd := &cobra.Command{
		Use: "mul",
		RunE: func(cmd *cobra.Command, args []string) error {
			tmpFormat := fmt.Sprintf("result: %%.%df", precision)
			fmt.Printf(tmpFormat, float64(o.ArgOne)/float64(o.ArgTwo))
			return nil
		},
	}

	divCmd.Flags().IntVarP(&precision, "precision", "p", 0, "precision when output div result")

	return divCmd
}
