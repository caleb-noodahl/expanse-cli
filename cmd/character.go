package cmd

import (
	"github.com/spf13/cobra"

	"github.com/caleb-noodahl/expanse-cli/engine"
)

// characterCMD represents the char command
var characterCMD = &cobra.Command{
	Use:   "char",
	Short: "character related actions",
	Long:  `generates, loads, updates, and rolls for a character`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.ParseFlags(args)

		action, _ := cmd.Flags().GetString("a")
		//name, _ := cmd.Flags().GetString("n")
		opts := map[string]func(*cobra.Command, []string){
			"gen":    engine.GenerateCharacter,
			"wiz":    engine.CharacterWizard,
			"load":   engine.LoadCharacter,
			"lvl":    engine.LevelUpCharacterWizard,
			"update": engine.UpdateCharacter,
		}

		f := opts[action]
		f(cmd, args)
	},
}

func init() {
	characterCMD.Flags().String("a", "", "action to route to:\n --a=gen\n --a=wiz\n --a=load\n --a=update\n --a=note")
	characterCMD.Flags().String("n", "", "name of character")
	rootCmd.AddCommand(characterCMD)
}
