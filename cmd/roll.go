package cmd

import (
	"fmt"

	"github.com/caleb-noodahl/expanse-cli/utils"
	"github.com/spf13/cobra"
)

// rollCmd represents the roll command
var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "simple dice roller",
	Long:  `rolls a dice for you based on --s[ize] and --a[mount] flags`,
	Run: func(cmd *cobra.Command, args []string) {
		sides, _ := cmd.Flags().GetInt("d")
		amount, _ := cmd.Flags().GetInt("a")
		for _, result := range utils.Roll(sides, amount) {
			fmt.Printf("1d%v: %v\n", sides, result)
		}
	},
}

func init() {
	rollCmd.Flags().Int("d", 6, "dice size")
	rollCmd.Flags().Int("a", 1, "amount to roll")
	rootCmd.AddCommand(rollCmd)
}
