package main

import (
	"borgmon"
	"borgmon/backend/hub"
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "borgmon/backend/migrations"
)

func main() {
	baseApp := getBaseApp()
	h := hub.NewHub(baseApp)
	if err := h.StartHub(); err != nil {
		log.Fatal(err)
	}
}

// getBaseApp creates a new PocketBase app with the default config
func getBaseApp() *pocketbase.PocketBase {
	isDev := os.Getenv("ENV") == "dev"

	baseApp := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: borgmon.AppName + "_data",
		DefaultDev:     isDev,
	})
	baseApp.RootCmd.Version = borgmon.Version
	baseApp.RootCmd.Use = borgmon.AppName
	baseApp.RootCmd.Short = ""
	// add update command
	// baseApp.RootCmd.AddCommand(&cobra.Command{
	// 	Use:   "update",
	// 	Short: "Update " + borgmon.AppName + " to the latest version",
	// 	Run:   hub.Update,
	// })
	// add health command
	// baseApp.RootCmd.AddCommand(newHealthCmd())

	// enable auto creation of migration files when making collection changes in the Admin UI
	migratecmd.MustRegister(baseApp, baseApp.RootCmd, migratecmd.Config{
		Automigrate: isDev,
		Dir:         "migrations",
	})

	return baseApp
}
