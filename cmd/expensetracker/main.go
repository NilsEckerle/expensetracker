package main

import (
	"log"

	"github.com/NilsEckerle/expensetracker/src/storage"
	"github.com/NilsEckerle/expensetracker/src/tui"
)

func main() {
	db, err := storage.Open("expenses.db")
	if err != nil {
		log.Fatalf("open db: %v", err)
	}

	// Example: seed a currency and create an expense.
	db.Create(&storage.Currency{CurrencyCode: "EUR", Symbol: "€", Ascii: "EUR"})

	_ = db // hand off to your TUI from here

	tui.Run()
}
