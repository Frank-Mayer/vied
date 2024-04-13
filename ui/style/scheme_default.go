//go:build !windows && !darwin && !linux

package style

func GetSystemColorScheme() (ColorScheme, error) {
	return ColorSchemeDefault, nil
}
