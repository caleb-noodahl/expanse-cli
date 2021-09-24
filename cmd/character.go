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
	"exp/functions"
	"fmt"

	"github.com/spf13/cobra"
)

// charCmd represents the char command
var charCmd = &cobra.Command{
	Use:   "char",
	Short: "character related actions",
	Long:  `generates, loads, updates, and rolls for a character`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.ParseFlags(args)
		action, _ := cmd.Flags().GetString("a")

		fmt.Printf("action:%s\n", action)
		switch action {
		case "gen":
			functions.GenerateCharacter(cmd, args)
		default:
			fmt.Println("unknown character action.")
		}
	},
}

func init() {
	charCmd.Flags().String("a", "", "character action to route to")
	charCmd.Flags().String("n", "", "character name")
	rootCmd.AddCommand(charCmd)
}
