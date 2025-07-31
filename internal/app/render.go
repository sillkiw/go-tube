package app

import (
	"gotube/internal/templates"
	"net/http"
)

func (app *Application) render(w http.ResponseWriter, name string, data any) {
	tmpl, ok := app.templateCache[name]
	if !ok {
		app.logger.Error("Template not found", "name", name)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		app.logger.Error("Template execution error", "name", name, "err", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}

func (app *Application) renderError(w http.ResponseWriter, msg string) {
	data := &templates.PageErr{ErrMsg: msg}
	app.render(w, "error.html", data)
}
