package functions

import "github.com/spf13/cobra"

type FunctionLoader interface {
	Generate() []*cobra.Command
}

type FunctionManager struct {
	Functions []FunctionLoader
	Commands  []func(*cobra.Command, []string)
}

func (f *FunctionManager) Load() {
	f.Commands = []func(*cobra.Command, []string){}

}
