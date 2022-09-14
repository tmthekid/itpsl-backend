package main

import "net/http"

func (app *Application) getExpenditureValues(w http.ResponseWriter, r *http.Request) {
	expenditureValues, err := app.models.GetExpenditureValues()
	if err != nil {
		app.errorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	app.toJSON(w, "values", expenditureValues, http.StatusOK)
}