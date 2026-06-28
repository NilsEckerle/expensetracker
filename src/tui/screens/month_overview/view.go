package monthoverview

import tea "charm.land/bubbletea/v2"

func (m Model) View() tea.View {
	v := tea.NewView(helpStyle.Render("Hello From Month Overview!"))
	v.AltScreen = true

	return v
}
