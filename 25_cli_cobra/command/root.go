package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cal",
	Short: "CAL is a very fast static site generator SHORT",
	Long:  `A Fast and Flexible Static Site Generator LONG`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("")
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
