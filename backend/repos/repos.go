package repos

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/pocketbase/pocketbase/core"
)

type RepoManager struct {
	app core.App
}

var borgExec = ""

func (rm *RepoManager) InitCLI() error {
	cmd := exec.Command("borg", "--version")

	if cmd.Err == nil {
		rm.app.Logger().Info("Using borg executable from $PATH")
		borgExec = "borg"
		return nil
	}

	lerr := cmd.Err.Error()
	if strings.Contains(lerr, "executable file not found") == false {
		return cmd.Err
	}

	localPath, err := os.Executable()
	if err != nil {
		return err
	}

	localPath = filepath.Dir(localPath)

	localPath = filepath.Join(localPath, "borg")
	cmd = exec.Command(localPath, "--version")
	if cmd.Err != nil {
		return err
	}

	borgExec = localPath

	return nil
}

func NewRepoManager(app core.App) *RepoManager {
	return &RepoManager{
		app: app,
	}
}

func (rm *RepoManager) CLI_Version() (string, error) {
	// Example:
	// borg --version
	// borg 1.2.8
	cmd := exec.Command(borgExec, "--version")
	outstream, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	versionstring := string(outstream)
	versionstring = strings.Replace(versionstring, "borg ", "", -1)
	versionstring = strings.TrimSpace(versionstring)

	return versionstring, nil
}

func (rm *RepoManager) CLI_RepoInfo_JSON(path string) (string, error) {
	info := ""
	cmd := exec.Command(borgExec, "info", "--json", path)
	outstream, err := cmd.CombinedOutput()
	if err != nil {
		return info, err
	}
	info = string(outstream)

	return info, nil
}

func (rm *RepoManager) CLI_RepoList_JSON(path string) (string, error) {
	info := ""
	cmd := exec.Command(borgExec, "list", "--json", path)
	outstream, err := cmd.CombinedOutput()
	if err != nil {
		return info, err
	}
	info = string(outstream)

	return info, nil
}

func (rm *RepoManager) UpdateBorgData_All() {
	rm.app.Logger().Info("Updating all records")
	recs, err := rm.app.FindAllRecords("repos")
	if err != nil {
		rm.app.Logger().Error("Error getting all repos", err.Error())
		return
	}

	rm.app.Logger().Info("Got repos to update", "count", len(recs))
	for _, rec := range recs {
		repo := rec.GetString("name")
		err = rm.UpdateBorgData(repo)
		if err != nil {
			rm.app.Logger().Error("Error updating repo", "repo", repo, "err", err.Error())
		}
	}
}

func (rm *RepoManager) UpdateBorgData(repoName string) error {
	rec, err := rm.app.FindFirstRecordByData("repos", "name", repoName)
	if err != nil {
		return err
	}

	borgInfo, err := rm.CLI_RepoInfo_JSON(rec.GetString("path"))
	if err != nil {
		return err
	}
	borgList, err := rm.CLI_RepoList_JSON(rec.GetString("path"))
	if err != nil {
		return err
	}

	rec.Set("borgInfo", borgInfo)
	rec.Set("borgList", borgList)
	rec.Set("borgUpdated", time.Now())
	err = rm.app.Save(rec)

	return err
}
