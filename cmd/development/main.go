package main

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/MGerasimchuk/bet-api/internal/adapter/repo/pg"
	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
	"github.com/MGerasimchuk/bet-api/internal/domain/service"
	"github.com/MGerasimchuk/bet-api/internal/infrastructure/config"
	"github.com/MGerasimchuk/bet-api/internal/usecase"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// for playing with components, run from root of repo:
// CONFIG_FILE=${PWD}/internal/infrastructure/config/config.yml GO111MODULE=on go run -mod vendor ${PWD}/cmd/development/main.go
func main() {
	cfg := config.GetRootConfig()

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		/*cfg.DB.Host*/ "localhost", cfg.DB.Port, cfg.DB.User, cfg.DB.Name, cfg.DB.Password,
	)
	db, err := gorm.Open("postgres", connectionString)
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
	_ = betUsecase

	transactionUsecase := usecase.NewTransactionUsecase(changeAccountBalanceService, transactionRepo)
	_ = transactionUsecase

	//fmt.Println(transactionUsecase.CancelTransaction("170f5049-6a13-4a8e-afdb-fd168471b90e"))

	fmt.Println(betUsecase.BetResultProcess("1bb63387-a1df-4158-96ec-8c6077f689ec", 100, uuid.New().String(), "win", "game"))
	fmt.Println(betUsecase.BetResultProcess("1bb63387-a1df-4158-96ec-8c6077f689ec", 100, uuid.New().String(), "win", "game"))
	fmt.Println(betUsecase.BetResultProcess("1bb63387-a1df-4158-96ec-8c6077f689ec", 300, uuid.New().String(), "lost", "game"))
	fmt.Println(betUsecase.BetResultProcess("1bb63387-a1df-4158-96ec-8c6077f689ec", 100, uuid.New().String(), "win", "game"))
	fmt.Println(betUsecase.BetResultProcess("1bb63387-a1df-4158-96ec-8c6077f689ec", 100, uuid.New().String(), "win", "game"))
}
