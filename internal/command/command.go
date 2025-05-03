package command

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type (
	CommandFunc func(cmd *cobra.Command, args []string) error

	CommandError struct {
		Code             string    `json:"code"`
		MessageToUser    any       `json:"message_to_user"`
		ErrorDescription any       `json:"error_description"`
		TraceId          uuid.UUID `json:"trace_id"`
	}
)

func (e CommandError) Error() string {
	jsonMessage, _ := json.MarshalIndent(e, "", " ")

	return string(jsonMessage)
}
