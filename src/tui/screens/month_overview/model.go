package monthoverview

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NilsEckerle/expensetracker/src/expensetracker/services/currency"
	"github.com/NilsEckerle/expensetracker/src/tui/components/display_units/month"
	"github.com/NilsEckerle/expensetracker/src/tui/themes"
)

func getHelpStyle() lipgloss.Style {
	theme := themes.GetActive()
	return lipgloss.NewStyle().
		Foreground(theme.MutedAlt).
		Background(theme.BackgroundAlt).
		Padding(0, 1)
}

func getAppStyle() lipgloss.Style {
	theme := themes.GetActive()
	return lipgloss.NewStyle().
		Background(theme.Background)
}

type Model struct {
	months []month.Model
	width  int
	height int
}

func NewModel() Model {
	var test_data []month.Model
	converter := currency.NewConverter(currency.NewCachingExchangeRateAPI(currency.FrankfurterExchangeRateAPI{}, time.Duration(6*time.Hour)))

	test_data = append(test_data, month.NewMonth(2026, time.January, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.February, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.March, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.April, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.May, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.June, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.July, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.August, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.September, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.October, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.November, *converter))
	test_data = append(test_data, month.NewMonth(2026, time.December, *converter))

	test_data = append(test_data, month.NewMonth(2027, time.January, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.February, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.March, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.April, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.May, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.June, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.July, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.August, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.September, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.October, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.November, *converter))
	test_data = append(test_data, month.NewMonth(2027, time.December, *converter))

	return Model{
		months: test_data,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
