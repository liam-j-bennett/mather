package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newSubtractCommand(client MatherClient) *cobra.Command {
	return &cobra.Command{
		Use:   "subtract",
		Short: "Subtract the second from the first of two (decimal) integers.",
		Long: `Subtract the second from the first of two (decimal) integers.
For example, "summer subtract 2 1" -> 1`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			lhn, rhn, err := validateNumbers(args[0], args[1])
			if err != nil {
				return err
			}
			result, err := client.Subtract(lhn, rhn)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "Numbers subtract to", result)
			return nil
		},
	}
}
