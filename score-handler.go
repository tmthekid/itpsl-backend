package main

import "net/http"

func (app *Application) getScores(w http.ResponseWriter, r *http.Request) {
	scores, err := app.models.GetScores()
	if err != nil {
		app.errorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	app.toJSON(w, "scores", scores, http.StatusOK)
}