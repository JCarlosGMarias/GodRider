// Package helpers provides a set of handful functions for presentation layer in terms of response formatting and data mapping
package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func SetCommonResponseEncoder(w *http.ResponseWriter) *json.Encoder {
	(*w).Header().Set("Content-Type", "application/json")
	return json.NewEncoder(*w)
}

func ParseBody(body io.ReadCloser, v interface{}) {
	reqBody, _ := ioutil.ReadAll(body)
	json.Unmarshal(reqBody, v)
}
