//go:build linux

package style

import (
	"os/exec"
	"strings"
)

func GetSystemColorScheme() (ColorScheme, error) {
	output, err := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "gtk-theme").Output()
	if err == nil {
		if strings.Contains(strings.ToLower(string(output)), "dark") {
			return ColorSchemeDark, nil
		} else {
			return ColorSchemeLight, nil
		}
	}
	return ColorSchemeDefault, err
}
