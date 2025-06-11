package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mhCmd = &cobra.Command{
	Use:   "mh",
	Short: "Mh SHORT",
	Long:  `Mh LONG`,
	Run: func(cmd *cobra.Command, args []string) {
		isFullName, _ := cmd.Flags().GetBool("fullname")

		if isFullName {
			fmt.Println("Mahadi Hassan Mithun")
		} else {
			fmt.Println("MH Mithun")
		}

	},
}

func init() {
	rootCmd.AddCommand(mhCmd)

	mhCmd.Flags().Bool("fullname", false, "First name, Last name, Nick nmae")
}
