package main

import "net/http"

func (app *Application) getExpenditures(w http.ResponseWriter, r *http.Request) {
	expenditures, err := app.models.GetExpenditures()
	if err != nil { 
		app.errorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	app.toJSON(w, "expenditures", expenditures, http.StatusOK)
}