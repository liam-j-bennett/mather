package commands

import (
	"github.com/spf13/cobra"
)

func NewMatherCommand(client MatherClient) *cobra.Command {
	command := &cobra.Command{
		Use:   "mather",
		Short: "Mather commands client for performing 'math' operations.",
		Long: `Mather commands client for performing 'math' operations.
The operation is defined by the second variable.
For example, "summer add 9 1", which returns 10`,
	}
	command.AddCommand(
		newAddCommand(client),
		newSubtractCommand(client),
		newMultiplyCommand(client),
		newDivideCommand(client),
	)
	return command
}
