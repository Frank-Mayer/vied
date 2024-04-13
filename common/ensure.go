package common

import (
	"os"
)

func EnsureDirExists(path string) error {
	// Check if the directory already exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Directory doesn't exist, create it
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		// Some error occurred other than the directory not existing
		return err
	}

	return nil
}
