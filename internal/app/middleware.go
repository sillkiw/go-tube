package app

import (
	"gotube/internal/cookie"
	"log/slog"
	"net/http"
	"slices"
)

// requireRole returns a middleware handler that enforces role-based access.
// If the `roles` slice is empty, access is allowed for everyone.
// Otherwise, it retrieves the user’s role from the cookie via cookie.UserRole.
// If no valid role is found or the role is not in `roles`, it logs an informational
// message and redirects the client to the login page.
func (app *Application) requireRole(roles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(roles) > 0 {
			role, err := cookie.UserRole(r)
			if err != nil || !slices.Contains(roles, role) {
				app.logger.Info("access denied",
					slog.String("role", role),
					slog.String("error", err.Error()),
				)
				http.Redirect(w, r, "/auth", http.StatusSeeOther)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// securedFileServer returns a file-serving handler for the given directory.
// It strips the provided URL prefix before serving files from the filesystem.
// Use this to expose static assets under a specific URL path.
func (app *Application) securedFileServer(prefix, dir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(prefix, http.FileServer(http.Dir(dir))).ServeHTTP(w, r)
	})
}
