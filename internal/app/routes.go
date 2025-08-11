package app

import (
	"net/http"
)

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle(
		"/upload",
		app.requireRole(
			app.config.Auth.UploadAllowedRoles,
			http.HandlerFunc(app.upload),
		),
	)
	mux.Handle(
		"/vp",
		app.requireRole(
			app.config.Auth.ViewAllowedRoles,
			http.HandlerFunc(app.videoPlayerHandler),
		),
	)
	mux.Handle(
		"/send",
		app.requireRole(
			app.config.Auth.UploadAllowedRoles,
			http.HandlerFunc(app.showUploadForm),
		),
	)
	mux.Handle(
		"/deleteVideo",
		app.requireRole(
			app.config.Auth.DeleteAllowedRoles,
			http.HandlerFunc(app.deleteVideo)),
	)
	mux.Handle(
		"/static/",
		app.requireRole(
			app.config.Auth.ViewAllowedRoles,
			app.securedFileServer("/static", app.config.UI.StaticPath),
		),
	)
	mux.Handle(
		"/converted/",
		app.requireRole(
			app.config.Auth.ViewAllowedRoles,
			app.securedFileServer("/converted", app.config.Video.ConvertPath),
		),
	)
	for _, path := range []string{"/", "/lst"} {
		mux.Handle(
			path,
			app.requireRole(
				app.config.Auth.ViewAllowedRoles,
				http.HandlerFunc(app.listFolderHandler),
			),
		)
	}
	mux.HandleFunc("/favicon.ico", http.HandlerFunc(app.faviconHandler))
	mux.HandleFunc("/queque", app.quequeSize)
	// mux.HandleFunc("/editconfig", editConfigHandler)
	// mux.HandleFunc("/save-config", saveConfigHandler)
	mux.HandleFunc("/auth", app.login)

	return mux

}
