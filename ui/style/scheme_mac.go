//go:build darwin

package style

import (
	"os/exec"
	"strings"
)

func GetSystemColorScheme() (ColorScheme, error) {
	output, err := exec.Command("defaults", "read", "-g", "AppleInterfaceStyle").Output()
	if err == nil {
		if strings.TrimSpace(string(output)) == "Dark" {
			return ColorSchemeDark, nil
		} else {
			return ColorSchemeLight, nil
		}
	}
	return ColorSchemeDefault, err
}
