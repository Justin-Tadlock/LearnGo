package main

import (
	"net/http"
)

func (app *Application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}
