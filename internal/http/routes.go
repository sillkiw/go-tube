package http

import (
	"github.com/go-chi/chi/v5/middleware"
	create "github.com/sillkiw/gotube/internal/http/api/create"
	mvLogger "github.com/sillkiw/gotube/internal/http/middleware"
)

func (s *Server) routes() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Recoverer)
	s.router.Use(mvLogger.New(s.l))
	s.router.Use(middleware.URLFormat)

	s.router.Post("/api/videos", create.New(s.l))
	// mux.Handle(
	// 	"/upload",
	// 	app.requireRole(
	// 		app.config.Auth.UploadAllowedRoles,
	// 		http.HandlerFunc(app.upload),
	// 	),
	// )
	// mux.Handle(
	// 	"/vp",
	// 	app.requireRole(
	// 		app.config.Auth.ViewAllowedRoles,
	// 		http.HandlerFunc(app.videoPlayerHandler),
	// 	),
	// )
	// mux.Handle(
	// 	"/send",
	// 	app.requireRole(
	// 		app.config.Auth.UploadAllowedRoles,
	// 		http.HandlerFunc(app.showUploadForm),
	// 	),
	// )
	// mux.Handle(
	// 	"/deleteVideo",
	// 	app.requireRole(
	// 		app.config.Auth.DeleteAllowedRoles,
	// 		http.HandlerFunc(app.deleteVideo)),
	// )
	// mux.Handle(
	// 	"/converted/",
	// 	app.requireRole(
	// 		app.config.Auth.ViewAllowedRoles,
	// 		app.securedFileServer("/converted", app.config.Video.ConvertPath),
	// 	),
	// )
	// for _, path := range []string{"/", "/lst"} {
	// 	mux.Handle(
	// 		path,
	// 		app.requireRole(
	// 			app.config.Auth.ViewAllowedRoles,
	// 			http.HandlerFunc(app.listFolderHandler),
	// 		),
	// 	)
	// }
	// mux.HandleFunc("/queque", app.quequeSize)
	// // mux.HandleFunc("/editconfig", editConfigHandler)
	// // mux.HandleFunc("/save-config", saveConfigHandler)
	// mux.HandleFunc("/auth", app.login)

}
