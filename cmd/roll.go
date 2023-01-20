/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
