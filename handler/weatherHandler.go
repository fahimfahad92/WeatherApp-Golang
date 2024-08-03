package handler

import (
	"WeatherApp/util"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

type WeatherHandler struct{}

func (weather *WeatherHandler) CurrentWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting current weather")

	baseUrl := viper.GetString("baseUrl")
	currentWeather := viper.GetString("currentWeather")
	apiKey := viper.GetString("apiKey")

	location, err := util.GetQueryParamAsString("location", r)
	if err != nil {
		util.HandleError(err, w)
		return
	}

	currentWeatherUrl := fmt.Sprintf("%s%s?key=%s&q=%s&aqi=yes",
		baseUrl, currentWeather, apiKey, location)

	response, err := http.Get(currentWeatherUrl)

	if err != nil {
		util.HandleError(err, w)
		return
	}

	if response.Status != "200 OK" {
		util.HandleResponseError(response, w)
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		util.HandleError(err, w)
		return
	}
	log.Println("Response found", len(responseData))
	w.Write(responseData)
}

func (weather *WeatherHandler) ForecastWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting weather forecast")

	baseUrl := viper.GetString("baseUrl")
	apiKey := viper.GetString("apiKey")
	forecastWeather := viper.GetString("forecastWeather")

	location, err := util.GetQueryParamAsString("location", r)
	if err != nil {
		util.HandleError(err, w)
		return
	}

	forecastUrl := fmt.Sprintf("%s%s?key=%s&q=%s&days=3&aqi=no&alerts=no",
		baseUrl, forecastWeather, apiKey, location)

	response, err := http.Get(forecastUrl)

	if err != nil {
		util.HandleError(err, w)
		return
	}

	if response.Status != "200 OK" {
		util.HandleResponseError(response, w)
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		util.HandleError(err, w)
		return
	}
	log.Println("Response found", len(responseData))
	w.Write(responseData)
}
