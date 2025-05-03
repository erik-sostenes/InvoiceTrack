package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

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

type ComputeInvoiceBalancesCommand struct{}

func NewComputeInvoiceBalancesCommand() (*ComputeInvoiceBalancesCommand, error) {
	return &ComputeInvoiceBalancesCommand{}, nil
}

func (c ComputeInvoiceBalancesCommand) Compute() CommandFunc {
	return func(cmd *cobra.Command, args []string) error {
		if err := cmd.ParseFlags(args); err != nil {
			return err
		}

		dir, err := cmd.Flags().GetString(dirFlag)
		if err != nil {
			return err
		}

		ext, err := cmd.Flags().GetString(extFlag)
		if err != nil {
			return err
		}

		cmd.Printf("Processing '%s' files in directory: %s\n", ext, dir)

		return nil
	}
}
