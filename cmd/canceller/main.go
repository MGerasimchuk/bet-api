package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	"github.com/MGerasimchuk/bet-api/internal/adapter/repo/pg"
	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
	"github.com/MGerasimchuk/bet-api/internal/domain/service"
	"github.com/MGerasimchuk/bet-api/internal/infrastructure/config"
	"github.com/MGerasimchuk/bet-api/internal/usecase"
)

const dateTimeLayout = "2006-01-02 15:04:05"

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

	transactionRepo := pg.NewTransactionRepository(db)
	transactionalChangeAccountBalanceRepositoryGenerator := func() repository.TransactionalChangeAccountBalanceRepository {
		transactionalDBInstance := db.Begin()
		accountRepo := pg.NewAccountRepository(transactionalDBInstance)
		transactionRepo := pg.NewTransactionRepository(transactionalDBInstance)

		return pg.NewChangeAccountBalanceRepo(transactionalDBInstance, accountRepo, transactionRepo)
	}
	changeAccountBalanceService := service.NewChangeAccountBalanceService(transactionalChangeAccountBalanceRepositoryGenerator)
	transactionUsecase := usecase.NewTransactionUsecase(changeAccountBalanceService, transactionRepo)
	_ = transactionUsecase

	ticker := time.NewTicker(time.Duration(cfg.Canceller.RepeatCancellationEveryMinutes * float64(time.Minute)))
	tickerDone := make(chan bool)

	go func() {
		for {
			select {
			case <-tickerDone:
				return
			case t := <-ticker.C:
				logger.Infof("[%s] Ticker triggered", t.Format(dateTimeLayout))

				wg := sync.WaitGroup{}
				transactions, err := transactionRepo.GetLast(cfg.Canceller.CancelTransactionsCount * 2)
				if err != nil {
					logger.Error(err)
				}

				for i := 1; i < len(transactions); i += 2 {
					wg.Add(1)

					go func(transactionId string) {
						defer wg.Done()
						err = transactionUsecase.CancelTransaction(transactionId)
						if err == nil {
							logger.Infof("[%s] Transaction with ID: %s, has been successfully canceled", time.Now().Format(dateTimeLayout), transactionId)
						} else {
							logger.Errorf("[%s] Transaction with ID: %s, hasn't been canceled due to error: %s", time.Now().Format(dateTimeLayout), transactionId, err.Error())
						}
					}(transactions[i].ID())
				}
				wg.Wait()
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ticker.Stop()
	tickerDone <- true

	logger.Info("Ticker stopped")

	if err := db.Close(); err != nil {
		logger.Fatal("Error during db connection close:", err)
	}

	logger.Info("Service stopped")
}
