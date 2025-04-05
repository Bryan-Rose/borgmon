package repos

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

type RepoManager struct {
	app core.App
}

func NewRepoManager(app core.App) *RepoManager {
	return &RepoManager{
		app: app,
	}
}

func (rm *RepoManager) CLI_Version() (string, error) {
	// borg --version
	// borg 1.2.8
	cmd := exec.Command("borg", "--version")
	outstream, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	versionstring := string(outstream)
	versionstring = strings.Replace(versionstring, "borg ", "", -1)
	versionstring = strings.TrimSpace(versionstring)

	return versionstring, nil
}

func (rm *RepoManager) CLI_RepoInfo() (BorgInfo, error) {
	info := BorgInfo{}
	cmd := exec.Command("borg", "info", "--json")
	outstream, err := cmd.CombinedOutput()
	if err != nil {
		return info, err
	}

	err = json.Unmarshal(outstream, &info)

	if err != nil {
		return info, err
	}

	return info, nil
}
