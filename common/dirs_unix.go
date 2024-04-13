//go:build darwin || linux

package common

import (
	"errors"
	"os"
	"path/filepath"
)

func getGlobalDirs() (dirs, error) {
	d := dirs{}

	home, homeErr := os.UserHomeDir()

	if xdgDataHome, ok := os.LookupEnv("XDG_DATA_HOME"); ok {
		d.Data = filepath.Join(xdgDataHome, "vied")
	} else if homeErr == nil {
		d.Data = filepath.Join(home, "vied", "data")
	} else {
		return d, errors.Join(errors.New("failed to get global data directory"), homeErr)
	}

	if xdgConfigHome, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		d.Config = filepath.Join(xdgConfigHome, "vied")
	} else if homeErr == nil {
		d.Config = filepath.Join(home, "vied", "config")
	} else {
		return d, errors.Join(errors.New("failed to get global config directory"), homeErr)
	}

	return d, nil
}
