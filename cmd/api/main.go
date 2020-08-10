package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/MGerasimchuk/bet-api/internal/adapter/controller"
	"github.com/MGerasimchuk/bet-api/internal/adapter/repo/pg"
	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
	"github.com/MGerasimchuk/bet-api/internal/domain/service"
	"github.com/MGerasimchuk/bet-api/internal/infrastructure/config"
	"github.com/MGerasimchuk/bet-api/internal/usecase"
	"github.com/MGerasimchuk/bet-api/pkg/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
)

func main() {
	cfg := config.GetRootConfig()

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.Level(cfg.Log.Level))

	dbConnectionString := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Name, cfg.DB.Password,
	)
	db, err := gorm.Open("postgres", dbConnectionString)
	if err != nil {
		panic(err)
	}

	accountRepo := pg.NewAccountRepository(db)
	transactionRepo := pg.NewTransactionRepository(db)
	transactionalChangeAccountBalanceRepositoryGenerator := func() repository.TransactionalChangeAccountBalanceRepository {
		transactionalDBInstance := db.Begin()
		accountRepo := pg.NewAccountRepository(transactionalDBInstance)
		transactionRepo := pg.NewTransactionRepository(transactionalDBInstance)

		return pg.NewChangeAccountBalanceRepo(transactionalDBInstance, accountRepo, transactionRepo)
	}
	changeAccountBalanceService := service.NewChangeAccountBalanceService(transactionalChangeAccountBalanceRepositoryGenerator)

	betUsecase := usecase.NewBetUsecase(changeAccountBalanceService, accountRepo, transactionRepo)

	gin.SetMode(cfg.HTTPServer.Mode)
	router := gin.Default()
	router.Use(ginlogrus.Logger(logger), gin.Recovery())
	router.Use(utils.RequestLogger(logger))

	betController := controller.NewBetController(betUsecase, logger)

	router.POST("/account/:accountId/bet", betController.BetPost)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.HTTPServer.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", err)
	}

	if err := db.Close(); err != nil {
		logger.Fatal("Error during db connection close:", err)
	}

	logger.Info("Service stopped")
}
