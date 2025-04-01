package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()
	app.RootCmd.Use = "Borgmon"

	// app.OnServe().BindFunc(func(se *core.ServeEvent) error {
	// 	// registers new "GET /hello" route
	// 	se.Router.GET("/hello", func(re *core.RequestEvent) error {
	// 		return re.String(200, "Hello world!")
	// 	})

	// 	return se.Next()
	// })

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
