package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/invoice-track/internal/dependency"
	"github.com/spf13/cobra"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	var cmd cobra.Command

	err := dependency.InjectCommand(ctx, &cmd)
	if err != nil {
		log.Fatal(err)
		return
	}

	cmd.Execute()
}
