package hub

import (
	"borgmon"
	"borgmon/backend/repos"
	"borgmon/frontend"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Hub struct {
	core.App
	// *alerts.AlertManager
	// um     *users.UserManager
	// rm     *records.RecordManager
	// sm     *systems.SystemManager
	rm          *repos.RepoManager
	pubKey      string
	appURL      string
	borgVersion string
}

func NewHub(app core.App) *Hub {
	hub := &Hub{}
	hub.App = app

	// hub.AlertManager = alerts.NewAlertManager(hub)
	// hub.um = users.NewUserManager(hub)
	// hub.rm = records.NewRecordManager(hub)
	// hub.sm = systems.NewSystemManager(hub)
	hub.rm = repos.NewRepoManager((hub))
	hub.appURL, _ = GetEnv("APP_URL")

	borgVersion, err := hub.rm.CLI_Version()
	if err != nil {
		println("Unable to retrieve borg cli version")
		println(err)
	} else {
		println("Borg version " + borgVersion)
	}

	hub.borgVersion = borgVersion
	return hub
}

// GetEnv retrieves an environment variable with a "BORGMON_HUB_" prefix, or falls back to the unprefixed key.
func GetEnv(key string) (value string, exists bool) {
	if value, exists = os.LookupEnv("BORGMON_HUB_" + key); exists {
		return value, exists
	}
	// Fallback to the old unprefixed key
	return os.LookupEnv(key)
}

func (h *Hub) StartHub() error {
	h.App.OnServe().BindFunc(func(e *core.ServeEvent) error {
		// initialize settings / collections
		if err := h.initialize(e); err != nil {
			return err
		}

		// sync systems with config
		// if err := syncSystemsWithConfig(e); err != nil {
		// 	return err
		// }

		// register api routes
		if err := h.registerApiRoutes(e); err != nil {
			return err
		}

		// register cron jobs
		if err := h.registerCronJobs(e); err != nil {
			return err
		}

		// start server
		if err := h.startServer(e); err != nil {
			return err
		}

		// // start system updates
		// if err := h.sm.Initialize(); err != nil {
		// 	return err
		// }

		return e.Next()
	})

	// TODO: move to users package
	// handle default values for user / user_settings creation
	// h.App.OnRecordCreate("users").BindFunc(h.um.InitializeUserRole)
	// h.App.OnRecordCreate("user_settings").BindFunc(h.um.InitializeUserSettings)

	if pb, ok := h.App.(*pocketbase.PocketBase); ok {
		// log.Println("Starting pocketbase")
		err := pb.Start()
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Hub) initialize(e *core.ServeEvent) error {
	return nil
}

// registerCronJobs sets up scheduled tasks
func (h *Hub) registerCronJobs(_ *core.ServeEvent) error {
	// delete old records once every hour
	// h.Cron().MustAdd("delete old records", "8 * * * *", h.rm.DeleteOldRecords)

	// create longer records every 10 minutes
	// h.Cron().MustAdd("create longer records", "*/10 * * * *", h.rm.CreateLongerRecords)
	return nil
}

// custom api routes
func (h *Hub) registerApiRoutes(se *core.ServeEvent) error {
	// // returns public key and version
	// se.Router.GET("/api/beszel/getkey", func(e *core.RequestEvent) error {
	// 	info, _ := e.RequestInfo()
	// 	if info.Auth == nil {
	// 		return apis.NewForbiddenError("Forbidden", nil)
	// 	}
	// 	return e.JSON(http.StatusOK, map[string]string{"key": h.pubKey, "v": beszel.Version})
	// })
	// // check if first time setup on login page
	// se.Router.GET("/api/beszel/first-run", func(e *core.RequestEvent) error {
	// 	total, err := h.CountRecords("users")
	// 	return e.JSON(http.StatusOK, map[string]bool{"firstRun": err == nil && total == 0})
	// })
	// // send test notification
	// se.Router.GET("/api/beszel/send-test-notification", h.SendTestNotification)
	// // API endpoint to get config.yml content
	// se.Router.GET("/api/beszel/config-yaml", h.getYamlConfig)
	// // create first user endpoint only needed if no users exist
	// if totalUsers, _ := h.CountRecords("users"); totalUsers == 0 {
	// 	se.Router.POST("/api/beszel/create-user", h.um.CreateFirstUser)
	// }
	return nil
}

// startServer starts the server for the Beszel (not PocketBase)
func (h *Hub) startServer(se *core.ServeEvent) error {
	// TODO: exclude dev server from production binary
	switch h.IsDev() {
	case true:
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   "localhost:5173",
		})
		se.Router.GET("/{path...}", func(e *core.RequestEvent) error {
			proxy.ServeHTTP(e.Response, e.Request)
			return nil
		})
	default:
		// parse app url
		parsedURL, err := url.Parse(h.appURL)
		if err != nil {
			return err
		}
		// fix base paths in html if using subpath
		basePath := strings.TrimSuffix(parsedURL.Path, "/") + "/"
		indexFile, _ := fs.ReadFile(frontend.DistDirFS, "index.html")
		indexContent := strings.ReplaceAll(string(indexFile), "./", basePath)
		indexContent = strings.Replace(indexContent, "{{V}}", borgmon.Version, 1)
		// set up static asset serving
		staticPaths := [2]string{"/static/", "/assets/"}
		serveStatic := apis.Static(frontend.DistDirFS, false)
		// get CSP configuration
		csp, cspExists := GetEnv("CSP")
		// add route
		se.Router.GET("/{path...}", func(e *core.RequestEvent) error {
			// serve static assets if path is in staticPaths
			for i := range staticPaths {
				if strings.Contains(e.Request.URL.Path, staticPaths[i]) {
					e.Response.Header().Set("Cache-Control", "public, max-age=2592000")
					return serveStatic(e)
				}
			}
			if cspExists {
				e.Response.Header().Del("X-Frame-Options")
				e.Response.Header().Set("Content-Security-Policy", csp)
			}
			return e.HTML(http.StatusOK, indexContent)
		})
	}
	return nil
}
