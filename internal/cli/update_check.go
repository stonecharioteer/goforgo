package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	updateCheckTimeout = 1200 * time.Millisecond
	updateInstallCmd   = "go install github.com/stonecharioteer/goforgo/cmd/goforgo@latest"
)

var defaultTagsURLForTests = "https://api.github.com/repos/stonecharioteer/goforgo/tags?per_page=1"
var updateCheckNow = time.Now
var updateCheckCacheTTL = 24 * time.Hour
var updateCachePathForTests string

var (
	updateNoticeMu sync.RWMutex
	updateNotice   string
)

type githubTag struct {
	Name string `json:"name"`
}

type semVersion struct {
	Major int
	Minor int
	Patch int
}

type updateCheckCache struct {
	LastChecked time.Time `json:"last_checked"`
	Current     string    `json:"current"`
	Latest      string    `json:"latest"`
	IsNewer     bool      `json:"is_newer"`
}

func maybeNotifyUpdate(w io.Writer, currentVersion string) {
	maybeNotifyUpdateWithConfig(w, currentVersion, http.DefaultClient, defaultTagsURLForTests)
}

func maybeNotifyUpdateWithConfig(w io.Writer, currentVersion string, client *http.Client, tagsURL string) {
	latest, isNewer, err := resolveUpdateStatus(currentVersion, tagsURL, client)
	if err != nil || !isNewer {
		setUpdateNotice("")
		return
	}

	setUpdateNotice(compactUpdateNotice(latest, currentVersion))
	if w == nil {
		return
	}
	_, _ = fmt.Fprint(w, cliUpdateNotice(latest, currentVersion))
}

func resolveUpdateStatus(currentVersion, tagsURL string, client *http.Client) (string, bool, error) {
	if cache, ok := loadFreshUpdateCache(currentVersion); ok {
		return cache.Latest, cache.IsNewer, nil
	}

	latest, isNewer, err := checkForUpdate(currentVersion, tagsURL, client)
	if err != nil {
		return "", false, err
	}

	_ = saveUpdateCache(updateCheckCache{
		LastChecked: updateCheckNow().UTC(),
		Current:     currentVersion,
		Latest:      latest,
		IsNewer:     isNewer,
	})

	return latest, isNewer, nil
}

func getCachedUpdateNotice() string {
	updateNoticeMu.RLock()
	defer updateNoticeMu.RUnlock()
	return updateNotice
}

func setUpdateNotice(notice string) {
	updateNoticeMu.Lock()
	defer updateNoticeMu.Unlock()
	updateNotice = notice
}

func cliUpdateNotice(latest, current string) string {
	return fmt.Sprintf("\n🔔 Update available: %s (current: %s)\n   Update with: %s\n\n", latest, current, updateInstallCmd)
}

func compactUpdateNotice(latest, current string) string {
	return fmt.Sprintf("Update available: %s (current: %s) • %s", latest, current, updateInstallCmd)
}

func checkForUpdate(currentVersion, tagsURL string, client *http.Client) (latest string, isNewer bool, err error) {
	if client == nil {
		client = http.DefaultClient
	}

	current, ok := parseSemVersion(currentVersion)
	if !ok {
		return "", false, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), updateCheckTimeout)
	defer cancel()

	tag, err := fetchLatestTag(ctx, client, tagsURL)
	if err != nil {
		return "", false, err
	}

	latestParsed, ok := parseSemVersion(tag)
	if !ok {
		return "", false, nil
	}

	return tag, current.lessThan(latestParsed), nil
}

func fetchLatestTag(ctx context.Context, client *http.Client, tagsURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tagsURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "goforgo-update-check")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("tags request returned status %d", resp.StatusCode)
	}

	var tags []githubTag
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return "", err
	}
	if len(tags) == 0 || tags[0].Name == "" {
		return "", fmt.Errorf("no tags returned")
	}

	return tags[0].Name, nil
}

func loadFreshUpdateCache(currentVersion string) (updateCheckCache, bool) {
	path, err := updateCachePath()
	if err != nil {
		return updateCheckCache{}, false
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return updateCheckCache{}, false
	}

	var cache updateCheckCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return updateCheckCache{}, false
	}

	if cache.Current != currentVersion {
		return updateCheckCache{}, false
	}

	if updateCheckNow().Sub(cache.LastChecked) > updateCheckCacheTTL {
		return updateCheckCache{}, false
	}

	if cache.IsNewer && cache.Latest == "" {
		return updateCheckCache{}, false
	}

	return cache, true
}

func saveUpdateCache(cache updateCheckCache) error {
	path, err := updateCachePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func updateCachePath() (string, error) {
	if updateCachePathForTests != "" {
		return updateCachePathForTests, nil
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(cacheDir, "goforgo", "update-check.json"), nil
}

func parseSemVersion(raw string) (semVersion, bool) {
	v := strings.TrimSpace(raw)
	v = strings.TrimPrefix(v, "v")
	if v == "" {
		return semVersion{}, false
	}

	// Drop any build metadata / prerelease suffix.
	if i := strings.IndexAny(v, "+-"); i >= 0 {
		v = v[:i]
	}

	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		return semVersion{}, false
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return semVersion{}, false
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return semVersion{}, false
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return semVersion{}, false
	}

	return semVersion{Major: major, Minor: minor, Patch: patch}, true
}

func (v semVersion) lessThan(other semVersion) bool {
	if v.Major != other.Major {
		return v.Major < other.Major
	}
	if v.Minor != other.Minor {
		return v.Minor < other.Minor
	}
	return v.Patch < other.Patch
}
