package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaksymSemk/WeatherForecastAPI/config"
	"github.com/MaksymSemk/WeatherForecastAPI/utils"
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/weather", h.getWeather).Methods("GET") // change to GET
}

func (h *Handler) getWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "Missing city query parameter")
		return
	}

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", config.Envs.WeatherForecastAPI, city)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		utils.WriteJSON(w, http.StatusBadGateway, "Failed to get weather data")
		return
	}
	defer resp.Body.Close()

	var weatherResponse struct {
		Current struct {
			TempC     float64 `json:"temp_c"`
			Humidity  int     `json:"humidity"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
		} `json:"current"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Error decoding response")
		return
	}

	result := map[string]interface{}{
		"temperature": weatherResponse.Current.TempC,
		"humidity":    weatherResponse.Current.Humidity,
		"description": weatherResponse.Current.Condition.Text,
	}

	utils.WriteJSON(w, http.StatusOK, result)
}
