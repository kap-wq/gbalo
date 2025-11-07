package main

import (
	"app/cmd/web/handler"
	"app/cmd/web/middleware"
	"app/internal"
	"app/internal/store"
	"log"
)

func initApp() *application {
	core, err := internal.Boot("web")
	if err != nil {
		log.Fatal(err)
	}

	myMiddleware := &middleware.Middleware{
		Core: core,
	}

	myHandlers := &handler.Handler{
		Core: core,
	}

	app := &application{
		Core:       core,
		Handler:    myHandlers,
		Middleware: myMiddleware,
	}

	app.Core.Routes.Mount("/", app.routes())
	store := store.NewStore(app.Core.DB.Pool)

	app.Middleware.Store = store
	myHandlers.Store = store
	return app
}
