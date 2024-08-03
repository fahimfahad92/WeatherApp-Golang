package application

import (
	"WeatherApp/config"
	"WeatherApp/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func loadRoutes(a *App) {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(config.JsonContentTypeHeader)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/weather", a.loadWeatherRoutes)

	a.router = router
}

func (a *App) loadWeatherRoutes(router chi.Router) {
	weatherHandler := &handler.WeatherHandler{}

	router.Get("/current", weatherHandler.CurrentWeather)
	router.Get("/forecast", weatherHandler.ForecastWeather)
}
