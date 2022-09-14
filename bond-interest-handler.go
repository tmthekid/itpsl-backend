package main

import "net/http"

func (app *Application) getBondsInterests(w http.ResponseWriter, r *http.Request) {
	bondInterests, err := app.models.GetBondInterests()
	if err != nil {
		app.errorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	app.toJSON(w, "interests", bondInterests, http.StatusOK)
}