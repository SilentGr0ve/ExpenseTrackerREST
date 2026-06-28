package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/config"
	logger "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/repository"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer cancel()

	appConfig := config.MustLoad()

	zapLogger, err := logger.NewLogger(appConfig.Logger)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer zapLogger.Close()

	zapLogger.Warn("logger initialized")

	pool, err := repository.NewPool(ctx, appConfig.Database)
	if err != nil {
		zapLogger.Fatal("failed to initialize postgres pool:", zap.Error(err))
	}
	defer pool.Close()

	zapLogger.Warn("postgres pool initialized")

}
