package main

import (
	"fmt"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/config"
	logger "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
)

func main() {
	appConfig := config.MustLoad()

	logger, err := logger.NewLogger(appConfig.Logger)
	if err != nil {
		fmt.Errorf("logger initialization: %w", err)
	}
	defer logger.Close()

	logger.Warn("logger initialized")
}
