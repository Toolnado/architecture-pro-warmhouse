package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"temterature-api/internal/models"
	"time"
)

type TemperatureHandler struct{}

func NewTemperatureHandler() *TemperatureHandler {
	return &TemperatureHandler{}
}

func (h *TemperatureHandler) InitRoutes(router *http.ServeMux) {
	router.HandleFunc("/temperature", h.GetTemperatureByLocation)
	router.HandleFunc("/temperature/{sensor_id}", h.GetSensorTemperature)

}

func (h *TemperatureHandler) GetTemperatureByLocation(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")

	response := models.TemperatureResponse{
		Value:     rand.Float64(),
		Timestamp: time.Now(),
		Location:  location,
		SensorID:  setSensorID(location),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *TemperatureHandler) GetSensorTemperature(w http.ResponseWriter, r *http.Request) {
	sensorID := r.PathValue("sensor_id")

	response := models.TemperatureResponse{
		Value:     rand.Float64(),
		Timestamp: time.Now(),
		Location:  setLocation(sensorID),
		SensorID:  sensorID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func setLocation(sensorID string) string {
	switch sensorID {
	case "1":
		return "Livin Room"
	case "2":
		return "Bedroom"
	case "3":
		return "Kitchen"
	default:
		return "Unknown"
	}
}

func setSensorID(location string) string {
	switch location {
	case "Living Room":
		return "1"
	case "Bedroom":
		return "2"
	case "Kitchen":
		return "3"
	default:
		return "0"
	}
}
