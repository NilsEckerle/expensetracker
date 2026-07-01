package main

import (
	"log"
	"os"
	"time"

	"github.com/NilsEckerle/expensetracker/src/storage"
	"github.com/NilsEckerle/expensetracker/src/tui"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func setupTestData() {
	db, err := storage.Open("expenses.db")
	if err != nil {
		log.Fatalf("open db: %v", err)
	}

	// Currency: insert only if missing (idempotent).
	currencies := []storage.Currency{
		{CurrencyCode: "EUR", Symbol: "€", Ascii: "EUR"},
		{CurrencyCode: "USD", Symbol: "$", Ascii: "USD"},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&currencies)

	// Expenses: only seed if the table is empty, so re-running doesn't duplicate.
	var count int64
	db.Model(&storage.Expense{}).Count(&count)
	if count > 0 {
		return
	}

	seedExpenses(db)
}

func seedExpenses(db *gorm.DB) {
	date := func(month time.Month, day int) time.Time {
		return time.Date(2026, month, day, 0, 0, 0, 0, time.UTC)
	}

	expenses := []storage.Expense{
		// Recurring monthly across the whole year (12 occurrences from January).
		{
			Title:         "Testing",
			Description:   "Testing December Rounding etc.",
			CurrencyCode:  "EUR",
			AmountCents:   10099,
			StartDate:     date(time.December, 1),
			Occurrences:   1,
			IntervalUnit:  "month",
			IntervalCount: 1,
		},
		{
			Title:         "Rent",
			Description:   "Monthly apartment rent",
			CurrencyCode:  "EUR",
			AmountCents:   95000,
			StartDate:     date(time.January, 1),
			Occurrences:   11,
			IntervalUnit:  "month",
			IntervalCount: 1,
		},
		{
			Title:         "Gym membership",
			Description:   "Monthly gym fee",
			CurrencyCode:  "EUR",
			AmountCents:   2999,
			StartDate:     date(time.January, 5),
			Occurrences:   11,
			IntervalUnit:  "month",
			IntervalCount: 1,
		},
		{
			Title:         "Streaming",
			Description:   "Video streaming subscription",
			CurrencyCode:  "EUR",
			AmountCents:   1299,
			StartDate:     date(time.January, 15),
			Occurrences:   11,
			IntervalUnit:  "month",
			IntervalCount: 1,
		},
		// Every two weeks.
		{
			Title:         "Groceries",
			Description:   "Biweekly grocery run",
			CurrencyCode:  "EUR",
			AmountCents:   6500,
			StartDate:     date(time.January, 3),
			Occurrences:   26,
			IntervalUnit:  "week",
			IntervalCount: 2,
		},
		// One-off expenses spread across the year.
		{
			Title:         "Flight booking",
			Description:   "Summer holiday flights",
			CurrencyCode:  "EUR",
			AmountCents:   42000,
			StartDate:     date(time.March, 12),
			Occurrences:   1,
			IntervalUnit:  "year",
			IntervalCount: 1,
		},
		{
			Title:         "New laptop",
			Description:   "Work laptop replacement",
			CurrencyCode:  "EUR",
			AmountCents:   139900,
			StartDate:     date(time.May, 20),
			Occurrences:   1,
			IntervalUnit:  "year",
			IntervalCount: 1,
		},
		{
			Title:         "Car insurance",
			Description:   "Annual car insurance premium",
			CurrencyCode:  "EUR",
			AmountCents:   78000,
			StartDate:     date(time.September, 1),
			Occurrences:   1,
			IntervalUnit:  "year",
			IntervalCount: 1,
		},
		{
			Title:         "Christmas gifts",
			Description:   "Presents for family",
			CurrencyCode:  "EUR",
			AmountCents:   30000,
			StartDate:     date(time.November, 10),
			Occurrences:   1,
			IntervalUnit:  "year",
			IntervalCount: 1,
		},
		// --- USD expenses ---
		{
			Title:         "Cloud hosting",
			Description:   "Monthly server hosting (USD billed)",
			CurrencyCode:  "USD",
			AmountCents:   2500,
			StartDate:     time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC),
			Occurrences:   12 * 10,
			IntervalUnit:  "month",
			IntervalCount: 1,
		},
		{
			Title:         "Cloud hosting",
			Description:   "Monthly server hosting (USD billed)",
			CurrencyCode:  "USD",
			AmountCents:   2500,
			StartDate:     time.Date(2027, time.February, 1, 0, 0, 0, 0, time.UTC),
			Occurrences:   3,
			IntervalUnit:  "month",
			IntervalCount: 1,
		},
		{
			Title:         "Cloud hosting",
			Description:   "Monthly server hosting (USD billed)",
			CurrencyCode:  "USD",
			AmountCents:   2500,
			StartDate:     time.Date(2027, time.February, 1, 0, 0, 0, 0, time.UTC),
			Occurrences:   6,
			IntervalUnit:  "week",
			IntervalCount: 1,
		},
		{
			Title:         "Software license",
			Description:   "Annual IDE subscription",
			CurrencyCode:  "USD",
			AmountCents:   19900,
			StartDate:     date(time.February, 1),
			Occurrences:   1,
			IntervalUnit:  "year",
			IntervalCount: 1,
		},
		{
			Title:         "Online course",
			Description:   "One-off course purchase",
			CurrencyCode:  "USD",
			AmountCents:   4999,
			StartDate:     date(time.June, 18),
			Occurrences:   1,
			IntervalUnit:  "year",
			IntervalCount: 1,
		},
		{
			Title:         "Conference ticket",
			Description:   "Tech conference admission",
			CurrencyCode:  "USD",
			AmountCents:   65000,
			StartDate:     date(time.October, 5),
			Occurrences:   1,
			IntervalUnit:  "year",
			IntervalCount: 1,
		},
	}

	if err := db.Create(&expenses).Error; err != nil {
		log.Fatalf("seed expenses: %v", err)
	}
}

func main() {
	// setup log
	logFile, err := os.OpenFile("expensetracker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatalf("Failed to open log file: %v\n", err)
	}
	log.SetOutput(logFile)

	setupTestData()
	tui.Run()
}
