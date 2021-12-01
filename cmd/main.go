package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lesson/adapter/repository"
	"github.com/lesson/usecase/transaction_process"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewTransactionRepositoryDb(db)
	usecase := transaction_process.NewTransactionProcess(repo)

	input := transaction_process.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    2000,
	}
	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	outputJSON, _ := json.Marshal(output)
	fmt.Println(string(outputJSON))
}
