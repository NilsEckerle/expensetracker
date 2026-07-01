package month

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NilsEckerle/expensetracker/src/expensetracker/services/currency"
	"github.com/NilsEckerle/expensetracker/src/expensetracker/services/database"
	"github.com/NilsEckerle/expensetracker/src/storage"
	"github.com/NilsEckerle/expensetracker/src/tui/themes"
)

const (
	borderWidth = 1
	paddingT    = 0
	paddingB    = 1
	paddingH    = 2
	textWidth   = 10
)

func getTitleStyle() lipgloss.Style {
	theme := themes.GetActive()
	style := lipgloss.NewStyle().Foreground(theme.Accent)
	return style
}

func getCurrencyStyle() lipgloss.Style {
	theme := themes.GetActive()
	style := lipgloss.NewStyle().Foreground(theme.Foreground)
	return style
}

func getMonthStyle() lipgloss.Style {
	theme := themes.GetActive()
	style := lipgloss.NewStyle().
		Padding(paddingT, paddingH, paddingB).
		Width(GetWidth()).
		Background(theme.Background).
		Border(lipgloss.RoundedBorder()).
		BorderBackground(theme.Background).
		BorderForeground(theme.Foreground)
	return style
}

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

// func (m Model) SumAmountCache() ([]CurrencySum, error) {
// 	db := database.Get()
// 	start := time.Date(m.year, m.month, 1, 0, 0, 0, 0, time.UTC)
// 	end := start.AddDate(0, 1, 0)
//
// 	var results []CurrencySum
// 	err := db.Model(&storage.Expense{}).
// 		Select("currency_code, COALESCE(SUM(amount_cents), 0) AS total").
// 		Where("start_date >= ? AND start_date < ?", start, end).
// 		Group("currency_code").
// 		Scan(&results).Error
// 	return results, err
// }

func (m Model) SumAmountCache() ([]CurrencySum, error) {
	db := database.Get()
	start := time.Date(m.year, m.month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	var expenses []storage.Expense
	if err := db.Where("start_date < ?", end).Find(&expenses).Error; err != nil {
		return nil, err
	}

	totals := make(map[string]int)
	for _, e := range expenses {
		count := occurrencesInRange(e, start, end)
		if count > 0 {
			totals[e.CurrencyCode] += e.AmountCents * count
		}
	}

	results := make([]CurrencySum, 0, len(totals))
	for code, total := range totals {
		results = append(results, CurrencySum{CurrencyCode: code, Total: total})
	}
	return results, nil
}

// occurrencesInRange returns how many occurrences of e fall within [start, end).
func occurrencesInRange(e storage.Expense, start, end time.Time) int {
	occDate := e.StartDate
	count := 0
	for i := 0; i < e.Occurrences; i++ {
		if !occDate.Before(end) {
			break // occurrences only move forward, no more chances
		}
		if !occDate.Before(start) {
			count++
		}
		occDate = addInterval(occDate, e.IntervalUnit, e.IntervalCount)
	}
	return count
}

func addInterval(t time.Time, unit string, count int) time.Time {
	switch unit {
	case "day":
		return t.AddDate(0, 0, count)
	case "week":
		return t.AddDate(0, 0, 7*count)
	case "month":
		return t.AddDate(0, count, 0)
	case "year":
		return t.AddDate(count, 0, 0)
	default:
		return t
	}
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

func GetWidth() int {
	return borderWidth + paddingH + textWidth + paddingH + borderWidth
}
