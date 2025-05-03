package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invoice-track/internal/business"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	dirFlag = "dir"
	extFlag = "ext"
)

func SetCommandComputeInvoiceBalances(cmdHandler *ComputeInvoiceBalancesCommand) (*cobra.Command, error) {
	if cmdHandler == nil {
		return nil, errors.New("missing 'ComputeInvoiceBalancesCommand' dependency")
	}

	cmd := &cobra.Command{
		Use:   "compute-invoice-balances",
		Short: "Compute total invoice balances from statement files",
		Run:   ErrorCommand(cmdHandler.Compute()),
	}

	flags := cmd.PersistentFlags()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	defaultXMLDir := filepath.Join(homeDir, "Documents", "xml_repo")

	flags.String(dirFlag, defaultXMLDir, "Path to the directory containing statement files")
	flags.String(extFlag, "xml", "File extension to filter (e.g., xml, json)")

	_ = cmd.MarkPersistentFlagDirname(dirFlag)
	_ = cmd.MarkPersistentFlagFilename(extFlag)

	return cmd, nil
}

type ComputeInvoiceBalancesCommand struct {
	calculator business.TotalAmountCalculator
}

func NewComputeInvoiceBalancesCommand(calculator business.TotalAmountCalculator) (*ComputeInvoiceBalancesCommand, error) {
	if calculator == nil {
		return nil, fmt.Errorf("missing %T dependency", calculator)
	}

	return &ComputeInvoiceBalancesCommand{
		calculator: calculator,
	}, nil
}

func (c ComputeInvoiceBalancesCommand) Compute() CommandFunc {
	return func(cmd *cobra.Command, args []string) error {
		if err := cmd.ParseFlags(args); err != nil {
			return err
		}

		dirPath, err := cmd.Flags().GetString(dirFlag)
		if err != nil {
			return err
		}

		extension, err := cmd.Flags().GetString(extFlag)
		if err != nil {
			return err
		}

		cmd.Printf("Processing '%s' files in directory: %s\n", extension, dirPath)

		totalAmountByECD, err := c.calculator.Calculate(dirPath, extension)
		if err != nil {
			return err
		}

		c.printTotalAmountSummary(totalAmountByECD)
		return nil
	}
}

func (c ComputeInvoiceBalancesCommand) printTotalAmountSummary(data business.Map) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"FUECD", "Invoice Type", "Total Amount"})

	for fuecd, invoices := range data {
		for invoiceType, total := range invoices {
			row := []string{
				fuecd,
				invoiceType,
				fmt.Sprintf("%.2f", total),
			}
			table.Append(row)
		}
	}

	table.Render()
}
