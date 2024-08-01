package application

import (
	"WeatherApp/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func loadRoutes(a *App) {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(jsonContentType)

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

func jsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
