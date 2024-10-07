package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"subtract", "minus"},
	Short:   "Subtract a number from another",
	Long:    "Carry out the difference between 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Subtracting %s from %s = %s", args[0], args[1], Subtact(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
