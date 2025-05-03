package dependency

import (
	"context"

	"github.com/invoice-track/internal/command"
	"github.com/spf13/cobra"
)

func InjectCommand(_ context.Context, cmd *cobra.Command) error {
	computeInvoiceBalances, err := command.NewComputeInvoiceBalancesCommand()
	if err != nil {
		return  err
	}

	setCommandComputeInvoiceBalances, err := command.SetCommandComputeInvoiceBalances(computeInvoiceBalances)
	if err != nil {
		return err
	}

	*cmd = cobra.Command{
		Use:   "Wholesale Electrical Market in Mexico",
		Short: "settlement account statements",
	}

	cmd.AddCommand(setCommandComputeInvoiceBalances)

	return nil
}
