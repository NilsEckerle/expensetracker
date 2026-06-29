package month

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NilsEckerle/expensetracker/src/expensetracker/services/currency"
	"github.com/NilsEckerle/expensetracker/src/expensetracker/services/database"
	"github.com/NilsEckerle/expensetracker/src/storage"
)

var (
	currencyStyle = lipgloss.NewStyle().Foreground(lipgloss.Red)
	monthStyle    = lipgloss.NewStyle().Padding(2, 2).Width(16)
)

type Model struct {
	year  int
	month time.Month

	amount_in_cent  int
	currency        storage.Currency
	displayCurrency string

	converter currency.Converter
}

type CurrencySum struct {
	CurrencyCode string
	Total        int
}

func (m Model) SumAmountCache() ([]CurrencySum, error) {
	db := database.Get()
	start := time.Date(m.year, m.month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	var results []CurrencySum
	err := db.Model(&storage.Expense{}).
		Select("currency_code, COALESCE(SUM(amount_cents), 0) AS total").
		Where("start_date >= ? AND start_date < ?", start, end).
		Group("currency_code").
		Scan(&results).Error
	return results, err
}

func NewMonth(year int, month time.Month, converter currency.Converter) Model {
	m := Model{
		year:            year,
		month:           month,
		converter:       converter,
		displayCurrency: "EUR", // TODO: Change to display currency
	}

	sums, err := m.SumAmountCache()
	if err != nil {
		// TODO: Handle error
		fmt.Printf("Error SumAmountCache: %v\n", err)
		return m
	}
	var total int
	for _, sum := range sums {
		v, err := m.converter.ConvertTo(m.displayCurrency, sum.Total, sum.CurrencyCode)
		if err != nil {
			// TODO: Handle error
			fmt.Printf("Error SumAmountCache: %v\n", err)
			return m
		}
		total += v
	}
	m.amount_in_cent = total
	err = database.Get().
		Where("currency_code = ?", m.displayCurrency).
		First(&m.currency).Error
	if err != nil {
		fmt.Printf("Error loading currency: %v\n", err)
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
