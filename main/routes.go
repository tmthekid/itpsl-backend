package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/api/v1/conditions", app.getConditions)
	router.HandlerFunc(http.MethodGet, "/api/v1/expenditures", app.getExpenditures)
	router.HandlerFunc(http.MethodGet, "/api/v1/scores", app.getScores)
	router.HandlerFunc(http.MethodPost, "/api/v1/results", app.saveResultHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/download", app.DownloadData)
	return app.enableCORS(router)
}

func (app *Application) toJSON(w http.ResponseWriter, wrap string, data interface{}, status int) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data
	json, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
	return nil
}

func (app *Application) errorJSON(w http.ResponseWriter, err string, status int) {
	type Error struct {
		Message string `json:"message"`
	}
	jsonError := Error{Message: err }
	app.toJSON(w, "error", jsonError, status)
}