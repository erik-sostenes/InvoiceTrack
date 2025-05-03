package command

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func ErrorCommand(commandFunc CommandFunc) func(cmd *cobra.Command, args []string) {
	if commandFunc == nil {
		panic("missing command func dependency")
	}

	return func(cmd *cobra.Command, args []string) {
		traceID := uuid.New()

		err := commandFunc(cmd, args)
		if err == nil {
			return
		}

		cmd.PrintErr(CommandError{
			Code:             "ERR_CMD_EXEC_FAIL",
			MessageToUser:    "An error occurred while processing your request.",
			ErrorDescription: "Command execution failed. Please check the parameters or try again.",
			TraceId:          traceID,
		})

	}
}
