package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_ "go.uber.org/automaxprocs"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := newJSONLogger(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return err
	}

	defer func() {
		err = logger.Sync()
	}()

	logger.Info("Hello world!", zap.String("location", "world"))

	return err
}

func newJSONLogger(level string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	if level != "" {
		var zapLevel zap.AtomicLevel
		err := zapLevel.UnmarshalText([]byte(level))
		if err != nil {
			return nil, err
		}
		config.Level = zapLevel
	}
	return config.Build()
}
