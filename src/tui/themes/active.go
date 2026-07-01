package themes

var activeTheme *Theme = &Default

func SetActive(t Theme) {
	activeTheme = &t
}

func GetActive() Theme {
	return *activeTheme
}
