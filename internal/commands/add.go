package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newAddCommand(client MatherClient) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add the two (decimal) integers together.",
		Long: `Add the two (decimal) integers together.
For example, "summer add 1 2" -> 3`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			lhn, rhn, err := validateNumbers(args[0], args[1])
			if err != nil {
				return err
			}
			result, err := client.Add(lhn, rhn)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "Numbers add up to", result)
			return nil
		},
	}
}
