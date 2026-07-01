package tui

import tea "charm.land/bubbletea/v2"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.activeScreen, cmd = m.activeScreen.Update(msg)
	return m, cmd
}
