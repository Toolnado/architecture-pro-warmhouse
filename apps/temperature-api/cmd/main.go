package main

import (
	"log"
	"net/http"
	"os"
	"temterature-api/internal/app"
	"temterature-api/internal/handlers"
	"temterature-api/internal/middlewares"
)

func main() {
	mux := http.NewServeMux()

	temperaturehandler := handlers.NewTemperatureHandler()
	temperaturehandler.InitRoutes(mux)

	addr := getEnv("PORT", ":8081")
	temperature := app.NewApp(addr, middlewares.LoggingMiddleware(mux))
	if err := temperature.Run(); err != nil {
		log.Println(err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
