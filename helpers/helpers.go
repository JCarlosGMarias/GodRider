// Package helpers provides a set of handful functions for presentation layer in terms of response formatting and data mapping
package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	allowedOrigin     = "http://127.0.0.1:8000"
	contentTypeHeader = "Content-Type"
	allowOriginHeader = "Access-Control-Allow-Origin"
)

func SetCommonResponseEncoder(w *http.ResponseWriter) *json.Encoder {
	(*w).Header().Set(contentTypeHeader, "application/json")
	(*w).Header().Set(allowOriginHeader, allowedOrigin)
	return json.NewEncoder(*w)
}

func SetOptionsResponseEncoder(w *http.ResponseWriter) {
	(*w).Header().Set(allowOriginHeader, allowedOrigin)
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, PUT, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", contentTypeHeader)
}

func ParseBody(body io.ReadCloser, v interface{}) {
	reqBody, _ := ioutil.ReadAll(body)
	json.Unmarshal(reqBody, v)
}
