package style

type ColorScheme uint8

const (
	ColorSchemeLight ColorScheme = iota
	ColorSchemeDark
	ColorSchemeDefault ColorScheme = ColorSchemeDark
)
