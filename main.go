package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger = logger.With("app", "agent")
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		exit(err, "godotenv.Load()")
	}

	client := anthropic.NewClient()
	tools := []ToolDefinition{ReadFileDefinition, ListFilesDefinition, EditFileDefinition}
	agent := NewAgent(&client, tools)
	err = agent.Run(ctx)
	if err != nil {
		exit(err, "agent.Run")
	}
}

func exit(err error, origin string) {
	slog.Error(origin, "error", err)
	os.Exit(1)
}
