package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/config"
	logger "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/repository"
	httpserver "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/middleware"
	auth_repository "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/auth/repository"
	auth_service "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/auth/service"
	auth_transport "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/auth/transport"
	categories_repository "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/categories/repository"
	categories_service "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/categories/service"
	categories_transport "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/categories/transport"
	expenses_repository "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/expenses/repository"
	expenses_service "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/expenses/service"
	expenses_transport "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/expenses/transport"
	statistics_repository "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/statistics/repository"
	statistics_service "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/statistics/service"
	statistics_transport "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/statistics/transport"
	web_repository "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/web/repository"
	web_service "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/web/service"
	web_transport "github.com/SilentGr0ve/ExpenseTrackerREST/internal/features/web/transport"
	"go.uber.org/zap"

	_ "github.com/SilentGr0ve/ExpenseTrackerREST/docs"
)

// @title 		Expense Tracker API
// @version 	1.0
// @description Expense Tracker REST-API scheme
// @host 		127.0.0.1:8080
// @basePath 	/api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer <token>

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

	authRepo := auth_repository.NewAuthRepository(pool, appConfig.Database.Timeout)
	authService := auth_service.NewAuthService(authRepo, appConfig.JWT.Secret, appConfig.JWT.AccessExpiry)
	authHandler := auth_transport.NewAuthHTTPHandler(authService)

	categoriesRepo := categories_repository.NewCategoriesRepository(pool, appConfig.Database.Timeout)
	categoriesService := categories_service.NewCategoryService(categoriesRepo)
	categoriesHandler := categories_transport.NewCategoriesHTTPHandler(categoriesService)

	expenseRepo := expenses_repository.NewExpensesRepository(pool, appConfig.Database.Timeout)
	expenseService := expenses_service.NewExpensesService(expenseRepo)
	expenseHandler := expenses_transport.NewExpensesHTTPHandler(expenseService)

	statisticsRepo := statistics_repository.NewStatisticsRepository(pool, appConfig.Database.Timeout)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepo)
	statisticsHandler := statistics_transport.NewStatisticsHTTPHandler(statisticsService)

	webRepo := web_repository.NewWebRepository()
	webService := web_service.NewWebService(webRepo)
	webHandler := web_transport.NewWebHTTPHandler(webService)

	httpServer := httpserver.NewHTTPServer(
		appConfig.Server,
		zapLogger,
		middleware.CORS(appConfig.Server.AllowedOrigins),
		middleware.RequestID(),
		middleware.Logger(zapLogger),
		middleware.Panic(),
	)

	publicRouterV1 := httpserver.NewAPIVersionRouter(
		1,
		nil,
	)
	publicRouterV1.AddRoutes(authHandler.Routes()...)

	protectedRouterV1 := httpserver.NewAPIVersionRouter(
		1,
		[]middleware.Middleware{middleware.Auth(appConfig.JWT.Secret)},
	)
	protectedRouterV1.AddRoutes(authHandler.ProtectedRoutes()...)
	protectedRouterV1.AddRoutes(categoriesHandler.ProtectedRoutes()...)
	protectedRouterV1.AddRoutes(expenseHandler.ProtectedRoutes()...)
	protectedRouterV1.AddRoutes(statisticsHandler.ProtectedRoutes()...)

	httpServer.RegisterAPIRouters(
		publicRouterV1,
		protectedRouterV1,
	)

	httpServer.RegisterSwagger()
	httpServer.RegisterWeb(webHandler.Routes()...)

	if err := httpServer.Run(ctx); err != nil {
		zapLogger.Error("HTTP server run error", zap.Error(err))
		os.Exit(1)
	}
}
