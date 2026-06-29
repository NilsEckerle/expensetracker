package month

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	var v tea.View
	v.SetContent(monthStyle.Render(
		currencyStyle.Render(
			fmt.Sprintf(
				"%.2f %s",
				float64(m.amount_in_cent)/100,
				m.currency.Symbol,
			),
		),
	))
	return v
}
