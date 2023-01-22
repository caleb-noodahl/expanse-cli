package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

var lookupCMD = &cobra.Command{
	Use:   "lookup",
	Short: "prints for word",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.ParseFlags(args)
		word, _ := cmd.Flags().GetString("w")
		fmt.Printf("word %s", word)
	},
}

func init() {
	lookupCMD.Flags().String("w", "word", "word to define")
	rootCmd.AddCommand(lookupCMD)
}
