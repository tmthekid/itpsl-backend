package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func (app *Application) saveResultHandler(w http.ResponseWriter, r *http.Request) {
	err := app.models.SaveResult(r)
	if err != nil {
		app.errorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	app.toJSON(w, "result", response("Results have been saved"), http.StatusOK)
}

func (app *Application) DownloadData(w http.ResponseWriter, r *http.Request){
	file, _ := app.models.Download()
	if file != nil {
		Openfile, _ := os.Open("data.xlsx")
		defer Openfile.Close()
		tempBuffer := make([]byte, 512)
		Openfile.Read(tempBuffer)
		FileContentType := http.DetectContentType(tempBuffer)
		FileStat, _ := Openfile.Stat()
		FileSize := strconv.FormatInt(FileStat.Size(), 10)
		Filename := "statistics"
		w.Header().Set("Content-Type", FileContentType+";"+Filename)
		w.Header().Set("Content-Length", FileSize)
		Openfile.Seek(0, 0)
		io.Copy(w, Openfile) 
	}
}