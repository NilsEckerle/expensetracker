package month

import (
	"fmt"
	"log"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	var v tea.View

	amount_in_eur := float64(m.amount_in_cent) / 100

	title := m.month.String()
	title = lipgloss.NewStyle().Width(textWidth).Align(lipgloss.Center).Render(title)
	title = getTitleStyle().Render(title)

	year := fmt.Sprintf("%d", m.year)
	year = lipgloss.NewStyle().Width(textWidth).Align(lipgloss.Center).Render(year)
	year = getTitleStyle().Render(year)

	amount := leftpad(fmt.Sprintf("%.2f", amount_in_eur), 8, " ")
	amount = fmt.Sprintf(
		"%s %s",
		amount,
		m.currency.Symbol,
	)
	amount = getCurrencyStyle().Render(amount)
	// log.Printf("amount in cents '%d', in euro '%f', in string '%s'\n", m.amount_in_cent, amount_in_eur, amount)

	v.SetContent(getMonthStyle().Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			year,
			"",
			amount,
		),
	))
	return v
}

func leftpad(s string, length int, char string) string {
	if length < len(s) {
		return s
	}
	var sb strings.Builder
	for range length - len(s) {
		sb.WriteString(char)
	}
	sb.WriteString(s)
	return sb.String()
}

func rightpad(s string, length int, char string) string {
	if length < len(s) {
		return s
	}
	var sb strings.Builder
	sb.WriteString(s)
	for range length - len(s) {
		sb.WriteString(char)
	}
	return sb.String()
}

func centerpad(s string, length int, char string) string {
	if strings.Contains(s, "\n") {
		log.Fatalf("centerpad failed. String '%s' contaisns a \\n.", s)
	}
	if length < len(s) {
		return s
	}
	left, right := 0, 0
	padWidth := (length - len(s))
	isEven := padWidth%2 == 0
	if !isEven {
		left += 1
		length -= 1
	}

	left += length / 2
	right += length / 2

	var sb strings.Builder
	for range left {
		sb.WriteString(char)
	}
	sb.WriteString(s)
	for range right {
		sb.WriteString(char)
	}
	return sb.String()
}
