package monthoverview

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	var v tea.View
	v.AltScreen = true

	// terminal width — get this from your model (set via tea.WindowSizeMsg)
	maxWidth := m.width

	var rows []string    // finished rows
	var current []string // boxes in the row being built
	currentWidth := 0

	for _, mo := range m.months {
		box := mo.View().Content // the rendered month string
		w := lipgloss.Width(box)

		// would adding this box overflow the row?
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

	grid := lipgloss.JoinVertical(lipgloss.Left, rows...)
	helpLine := helpStyle.Render("q: quit")

	v.SetContent(lipgloss.JoinVertical(lipgloss.Left, grid, helpLine))
	return v
}
