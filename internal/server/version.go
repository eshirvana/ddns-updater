package server

import (
	"encoding/json"
	"net/http"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
}

func getLatestRelease() (string, error) {
	resp, err := http.Get("https://api.github.com/repos/qdm12/ddns-updater/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	return release.TagName, nil
}

// Add a function to get the current version
func getCurrentVersion() string {
	// Replace with actual logic to get the current version
	return "v2.9.0"
}
