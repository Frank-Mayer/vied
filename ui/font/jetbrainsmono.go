package font

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Frank-Mayer/vied/common"
	"github.com/Frank-Mayer/vied/api"
)

func getJetbrainsMono() (string, error) {
	// check if already available
	fontFileLocation := filepath.Join(common.Dirs.Data, "jetbrains-mono-nerd.ttf")
	if exists, err := common.Exists(fontFileLocation); exists {
		return fontFileLocation, nil
	} else if err != nil {
		return "", errors.Join(fmt.Errorf("failed to check for %s", fontFileLocation), err)
	}

	// download
	if rel, err := api.GetLatestRelease("ryanoasis", "nerd-fonts"); err != nil {
		return "", errors.Join(errors.New("failed to get latest release for ryanoasis/nerd-fonts"), err)
	} else {
		for _, asset := range rel.Assets {
			if asset.ContentType != "application/zip" {
				continue
			}
			if !strings.Contains(asset.Name, "JetBrains") {
				continue
			}

		}
	}

	return fontFileLocation, nil
}
