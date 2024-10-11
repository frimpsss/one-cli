package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shouldRoundUp bool
var multiplyCmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"time", "by"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Multiplying %s by %s is %s ", args[0], args[1], Multiply(args[0], args[1], shouldRoundUp))
	},
}

func init() {
	multiplyCmd.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round to 2 dp")
	rootCmd.AddCommand(multiplyCmd)
}
