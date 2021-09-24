package functions

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GenerateCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("generating character: %s\n", args)
	
}
