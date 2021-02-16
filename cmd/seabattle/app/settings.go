package app

type settings struct {
	theme Theme
}

func (s *settings) Theme() Theme {
	return s.theme
}

func (s *settings) SetTheme(theme Theme) {
	s.theme = theme
}
