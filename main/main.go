package main

import (
	"context"
	"itpsl/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

type Application struct {
	logger *log.Logger
	models models.Connection
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
	dsn := "postgres://"+os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_NAME")
	db, err := openDB(dsn)
	if err != nil {
		log.Fatal("Database connection failed")
	}
	defer db.Close()
	app := &Application {
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		models: models.GetModels(db),
	}
	server := &http.Server {
		Addr: ":"+os.Getenv("SERVER_PORT"),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.logger.Println("Server is running on port "+os.Getenv("SERVER_PORT"))
	err = server.ListenAndServe()
	if err != nil {
		app.logger.Fatal("Server connection failed")
	}
}

func openDB(dsn string)(*pgxpool.Pool, error){
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}