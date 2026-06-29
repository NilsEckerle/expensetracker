package database

import (
	"fmt"

	"github.com/NilsEckerle/expensetracker/src/storage"
	"gorm.io/gorm"
)

func Get() *gorm.DB {
	db, err := storage.Open("expenses.db")
	if err != nil {
		// TODO: Handle error
		fmt.Printf("Error opening database: %v\n", err)
		return nil
	}

	return db
}
