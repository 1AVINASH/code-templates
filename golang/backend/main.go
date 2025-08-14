package main

import (
	"log"
	"net/http"
	"os"

	"gotemplate/infra/postgres"
	redisclient "gotemplate/infra/redis"
	logger "gotemplate/utility/logger"

	user "gotemplate/services/users"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Set up logger
	logger, err := logger.New("debug")
	if err != nil {
		log.Fatalf("Failed to init logger %v", err)
	}
	logger.Info("Starting server")

	// Set up Environment variables
	err = godotenv.Load("dev.env")
	if err != nil {
		log.Println("No .env file found or unable to load it:", err)
	}

	// Setting App Mode
	appMode := os.Getenv("APP_MODE")
	logger.Infof("App Mode is %s", appMode)

	if appMode == "API" {
		mux := chi.NewRouter()

		// Set up Infrastructural components
		redisclient.Initialize()
		postgres.Inititialize()

		defer postgres.ClosePostgres()
		defer redisclient.CloseRedis()

		// Add all services here
		user.NewUserManager(mux)

		// Start the server and listen on port 8080.
		port := ":8080"
		log.Fatal(http.ListenAndServe(port, mux))

	} else if appMode == "CRON" {
		//
	}
}
