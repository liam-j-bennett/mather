package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func newDivideCommand(client MatherClient) *cobra.Command {
	return &cobra.Command{
		Use:   "divide",
		Short: "Divide the first by the second of two (decimal) integers to 5DP.",
		Long: `Divide the first by the second of two (decimal) integers to 5DP.
For example, "summer divide 1 2" -> 0.5`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			lhn, rhn, err := validateNumbers(args[0], args[1])
			if err != nil {
				return err
			}
			result, err := client.Divide(lhn, rhn)
			if err != nil {
				return err
			}
			stringResult := strconv.FormatFloat(result, 'g', 5, 64)
			fmt.Fprintln(cmd.OutOrStdout(), "Numbers divide to", stringResult)
			return nil
		},
	}
}
