package main

import (
	"net/http"
)

func (app *Application) getConditions(w http.ResponseWriter, r *http.Request) {
	conditions, err := app.models.GetConditions()
	if err != nil {
		app.errorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	app.toJSON(w, "conditions", conditions, http.StatusOK)
}