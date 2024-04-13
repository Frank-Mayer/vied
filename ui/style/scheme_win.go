//go:build windows

package style

import (
	"os/exec"
	"strings"
)

func GetSystemColorScheme() (ColorScheme, error) {
	output, err := exec.Command("reg", "query", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize", "/v", "AppsUseLightTheme").Output()
	if err == nil {
		if strings.Contains(string(output), "0x1") {
			return ColorSchemeDark, nil
		} else {
			return ColorSchemeLight, nil
		}
	}
	return ColorSchemeDefault, err
}
