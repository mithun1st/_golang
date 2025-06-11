package command

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

var piCmd = &cobra.Command{
	Use:   "pi",
	Short: "PI SHORT",
	Long:  `PI LONG`,
	Run: func(cmd *cobra.Command, args []string) {

		plusNum, _ := cmd.Flags().GetInt("plus")
		isShort, _ := cmd.Flags().GetBool("short")
		title, _ := cmd.Flags().GetString("title")

		if plusNum != 0 {
			fmt.Println(math.Pi + float64(plusNum))
		} else if isShort {
			fmt.Printf("%.2f\n", math.Pi)
		} else {
			fmt.Println(title, ":", math.Pi)
		}

	},
}

func init() {
	rootCmd.AddCommand(piCmd)

	piCmd.Flags().IntP("plus", "p", 0, "Plus number with pi")
	piCmd.Flags().BoolP("short", "s", false, "View in short form")
	piCmd.Flags().String("title", "PI Value", "Title of pi")
}
