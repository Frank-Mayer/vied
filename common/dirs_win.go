//go:build windows

package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func getGlobalDirs() (dirs, error) {
	d := dirs{}

	home, homeErr := os.UserHomeDir()

	if appdata, ok := os.LookupEnv("APPDATA"); ok {
		d.Data = filepath.Join(appdata, "vied")
	} else if homeErr == nil {
		d.Data = filepath.Join(home, "vied", "data")
	} else {
		return d, errors.Join(errors.New("failed to get global data directory"), homeErr)
	}

	if err := EnsureDirExists(d.Data); err != nil {
		return d, errors.Join(fmt.Errorf("failed to create data directory %s", d.Data), err)
	}

	if localAppdata, ok := os.LookupEnv("LOCALAPPDATA"); ok {
		d.Config = filepath.Join(localAppdata, "vied")
	} else if homeErr == nil {
		d.Config = filepath.Join(home, "vied", "config")
	} else {
		return d, errors.Join(errors.New("failed to get global config directory"), homeErr)
	}

	if err := EnsureDirExists(d.Config); err != nil {
		return d, errors.Join(fmt.Errorf("failed to create config directory %s", d.Data), err)
	}

	return d, nil
}
