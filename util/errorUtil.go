package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type WeatherErrorResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func HandleError(err error, w http.ResponseWriter) http.ResponseWriter {
	log.Println(err.Error())
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
	return w
}

func HandleResponseError(response *http.Response, w http.ResponseWriter) http.ResponseWriter {
	var weatherErrorResponse WeatherErrorResponse

	if err := json.NewDecoder(response.Body).Decode(&weatherErrorResponse); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return w
	}

	var errorResponse ErrorResponse
	errorResponse.ErrorMessage = weatherErrorResponse.Error.Message

	log.Printf("response error %s\n", errorResponse.ErrorMessage)

	res, err := json.Marshal(errorResponse)
	if err != nil {
		log.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return w
	}

	w.WriteHeader(response.StatusCode)
	w.Write(res)

	return w
}
