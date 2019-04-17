package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newMultiplyCommand(client MatherClient) *cobra.Command {
	return &cobra.Command{
		Use:   "multiply",
		Short: "Multiply the two (decimal) integers together.",
		Long: `Multiply the two (decimal) integers together.
For example, "summer multiply 2 1" -> 2`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			lhn, rhn, err := validateNumbers(args[0], args[1])
			if err != nil {
				return err
			}
			result, err := client.Multiply(lhn, rhn)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "Numbers multiply to", result)
			return nil
		},
	}
}
