package themes

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Theme is a named palette. Add fields as your UI needs more roles.
type Theme struct {
	Name          string
	Background    color.Color // base backdrop
	BackgroundAlt color.Color // raised surfaces (boxes, panels)
	Foreground    color.Color
	Accent        color.Color // e.g. currency text
	Muted         color.Color // secondary text on Background
	MutedAlt      color.Color // secondary text on BackgroundAlt
	Border        color.Color
}

var Default = Nord

var Gruvbox = Theme{
	Name:          "gruvbox",
	Background:    lipgloss.Color("#1d2021"),
	BackgroundAlt: lipgloss.Color("#282828"),
	Foreground:    lipgloss.Color("#ebdbb2"),
	Accent:        lipgloss.Color("#fb4934"),
	Muted:         lipgloss.Color("#928374"), // gruvbox gray, reads on base
	MutedAlt:      lipgloss.Color("#a89984"), // lighter gray for the surface
	Border:        lipgloss.Color("#665c54"),
}

var Dracula = Theme{
	Name:          "dracula",
	Background:    lipgloss.Color("#21222c"),
	BackgroundAlt: lipgloss.Color("#282a36"),
	Foreground:    lipgloss.Color("#f8f8f2"),
	Accent:        lipgloss.Color("#ff79c6"),
	Muted:         lipgloss.Color("#6272a4"), // dracula comment color
	MutedAlt:      lipgloss.Color("#7b89bd"), // lightened so it lifts off the surface
	Border:        lipgloss.Color("#44475a"),
}

var Nord = Theme{
	Name:          "nord",
	Background:    lipgloss.Color("#2e3440"),
	BackgroundAlt: lipgloss.Color("#3b4252"),
	Foreground:    lipgloss.Color("#d8dee9"),
	Accent:        lipgloss.Color("#88c0d0"),
	Muted:         lipgloss.Color("#4c566a"), // nord3, reads on nord0 base
	MutedAlt:      lipgloss.Color("#7b88a1"), // lightened nord3 for nord1 surface
	Border:        lipgloss.Color("#434c5e"),
}

var All = []Theme{Default, Gruvbox, Dracula, Nord}

func Next(current Theme) Theme {
	for i, t := range All {
		if t.Name == current.Name {
			return All[(i+1)%len(All)]
		}
	}
	return All[0]
}
