package main

import (
	"database/sql"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/event"
	createaccount "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_account"
	createclient "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_client"
	createtransaction "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com.br/devfullcycle/fc-ms-wallet/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)	

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDb, accountDb,eventDispatcher, transactionCreatedEvent)
}