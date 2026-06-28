package storage

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Open(path string) (*gorm.DB, error) {
	// DisableForeignKeyConstraintWhenMigrating stops GORM from emitting
	// physical FOREIGN KEY clauses in the CREATE TABLE statements. The
	// associations still work at the ORM level (Preload, joins), but GORM
	// no longer auto-generates FK constraints, which is what it was
	// getting wrong (it inverted the Expense->Currency relationship and
	// put an integer FK on currencies pointing back at expenses).
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Currency{}, &Tag{}, &Expense{}); err != nil {
		return nil, err
	}

	return db, nil
}
