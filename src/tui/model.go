package tui

import (
	tea "charm.land/bubbletea/v2"
	monthoverview "github.com/NilsEckerle/expensetracker/src/tui/screens/month_overview"
)

type model struct {
	activeScreen tea.Model
	width        int
	height       int
}

func NewModel() model {
	return model{
		activeScreen: monthoverview.NewModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
