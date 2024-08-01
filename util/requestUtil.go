package util

import (
	"net/http"
)

func GetQueryParamAsString(param string, r *http.Request) (string, error) {
	return r.URL.Query().Get(param), nil
}
