package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type (
	Release struct {
		TagName         string  `json:"tag_name"`
		TargetCommitish string  `json:"target_commitish"`
		Name            string  `json:"name"`
		Draft           bool    `json:"draft"`
		Prerelease      bool    `json:"prerelease"`
		Assets          []Asset `json:"assets"`
	}

	Asset struct {
		Name               string `json:"name"`
		ContentType        string `json:"content_type"`
		Size               int    `json:"size"`
		BrowserDownloadUrl string `json:"browser_download_url"`
	}
)

func GetLatestRelease(owner string, repo string) (Release, error) {
	r := Release{}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return r, errors.Join(fmt.Errorf("failed to request %s", url), err)
	}
	if resp.StatusCode >= 300 {
		return r, fmt.Errorf("failed to request %s: %s", url, resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return r, errors.Join(errors.New("failed to read json from response body"), err)
	}

	return r, nil
}
