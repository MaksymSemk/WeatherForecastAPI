package weather

import (
	"net/http"

	"github.com/MaksymSemk/WeatherForecastAPI/utils"
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/weather", h.getWeather).Methods("GET")
}

func (h *Handler) getWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "Missing 'city' query parameter", http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"temperature":   21.5,
		"humidity":      65,
		"description":   "Partly cloudy",
		"requestedCity": city,
	})
}
