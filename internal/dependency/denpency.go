package dependency

import (
	"context"

	"github.com/invoice-track/internal/business"
	"github.com/invoice-track/internal/command"
	"github.com/spf13/cobra"
)

func InjectCommand(_ context.Context, cmd *cobra.Command) error {
	totalAmountCalculator, err := business.NewTotalAmountCalculator()
	if err != nil {
		return err
	}

	computeInvoiceBalances, err := command.NewComputeInvoiceBalancesCommand(totalAmountCalculator)
	if err != nil {
		return err
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
