package repos

import "github.com/pocketbase/pocketbase/core"

type RepoManager struct {
	app core.App
}

func NewRepoManager(app core.App) *RepoManager {
	return &RepoManager{
		app: app,
	}
}
