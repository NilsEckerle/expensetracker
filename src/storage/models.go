package storage

import "time"

// Currency is a supported currency, referenced by expenses via CurrencyCode.
type Currency struct {
	CurrencyCode string `gorm:"primaryKey;type:varchar(10)"`
	Symbol       string `gorm:"type:varchar(4);not null"`
	Ascii        string `gorm:"type:varchar(32);not null"`
}

// Tag is a label that can be attached to any number of expenses.
type Tag struct {
	TagName     string `gorm:"primaryKey;type:varchar(64)"`
	Description string `gorm:"not null"`

	Expenses []Expense `gorm:"many2many:expenses_tags;"`
}

// Expense is one (possibly recurring) expense definition.
type Expense struct {
	ExpenseID     uint      `gorm:"primaryKey;autoIncrement"`
	Title         string    `gorm:"not null"`
	Description   string    `gorm:"not null"`
	CurrencyCode  string    `gorm:"type:varchar(10);not null"`
	AmountCents   int       `gorm:"not null"`
	StartDate     time.Time `gorm:"not null"`
	Occurrences   int       `gorm:"not null"`
	IntervalUnit  string    `gorm:"not null"` // day | week | month | year
	IntervalCount int       `gorm:"not null"`

	// Belongs-to. The FK column CurrencyCode lives on THIS struct and
	// references Currency's primary key. Both foreignKey and references
	// are given explicitly so GORM does not infer the relationship in the
	// wrong direction (which made currencies hold an integer FK back to
	// expenses).
	Currency Currency `gorm:"foreignKey:CurrencyCode;references:CurrencyCode"`

	// Many-to-many through the expenses_tags join table.
	Tags []Tag `gorm:"many2many:expenses_tags;"`
}
