package monthoverview

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NilsEckerle/expensetracker/src/tui/components/display_units/month"
)

func (m Model) buildWrappedMonthRows(width int) []string {
	maxWidth := width

	var rows []string    // finished rows
	var current []string // boxes in the row being built
	currentWidth := 0

	for _, mo := range m.months {
		box := mo.View().Content
		w := lipgloss.Width(box)

		if currentWidth+w > maxWidth && len(current) > 0 {
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, current...))
			current = nil
			currentWidth = 0
		}

		current = append(current, box)
		currentWidth += w
	}
	if len(current) > 0 {
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, current...))
	}

	var paddedRows []string
	for _, row := range rows {
		paddedRows = append(paddedRows, getAppStyle().Width(m.width).Align(lipgloss.Center).Render(row))
	}

	return paddedRows
}

func (m Model) View() tea.View {
	var v tea.View
	v.AltScreen = true

	paddedRows := m.buildWrappedMonthRows(month.GetWidth() * 4)
	grid := lipgloss.JoinVertical(lipgloss.Left, paddedRows...)
	helpLine := getHelpStyle().Width(m.width).Render("q: quit")

	// Height available for the grid = full height minus the help line.
	gridHeight := m.height - lipgloss.Height(helpLine)

	// Place the grid in the top region...
	top := lipgloss.Place(
		m.width, gridHeight,
		lipgloss.Left, lipgloss.Top,
		grid,
		lipgloss.WithWhitespaceStyle(getAppStyle()),
	)

	// ...then stack the help line under it, sitting at the bottom.
	content := lipgloss.JoinVertical(lipgloss.Left, top, helpLine)

	v.SetContent(content)
	return v
}
