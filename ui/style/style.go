package style

import (
	"image/color"

	"github.com/charmbracelet/log"
)

var (
	Background color.RGBA
)

func init() {
	scheme, err := GetSystemColorScheme()
	if err != nil {
		log.Error("failed to get system color scheme", "err", err)
	}
	switch scheme {
	case ColorSchemeLight:
		Background = color.RGBA{239, 241, 245, 255}

	case ColorSchemeDark:
		Background = color.RGBA{30, 30, 46, 255}

	default:
		log.Fatal("invalid color scheme", "code", scheme)
	}
}
