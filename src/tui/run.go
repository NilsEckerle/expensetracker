package tui

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func Run() {
	if _, err := tea.NewProgram(NewModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
