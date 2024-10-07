package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "one",
	Short: "one is a cli to perform basic math ops",
	Long:  "one is a cli to perform basic math ops",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, "Opps, an error occurred while execution")
		os.Exit(1)
	}
}
