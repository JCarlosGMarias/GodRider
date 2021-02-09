// helpers package provides a set of handful functions for presentation layer in terms of requests validation and data mapping
package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"godrider/dtos/responses"
	"godrider/services"
)

// IsValidMethod reads the Request's Method field, searches it into allowedMethods and determines
// if is an allowed method for the calling controller endpoint
func IsValidMethod(w http.ResponseWriter, r *http.Request, allowedMethods []string) bool {
	isValidMethod := false
	for _, m := range allowedMethods {
		if r.Method == m {
			isValidMethod = true
			break
		}
	}

	if !isValidMethod {
		message := fmt.Sprintf("Incoming request with method %s not allowed!", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Code: responses.METHOD_UNAUTHORIZED, Message: message})
	}
	return isValidMethod
}

// IsValidToken determine if the request token belong to an existing or active user
func IsValidToken(w http.ResponseWriter, r *http.Request, token string) bool {
	isValidToken := false
	for _, user := range services.UserSrv.GetAllUsers() {
		if user.Token == token {
			isValidToken = true
			break
		}
	}

	if !isValidToken {
		json.NewEncoder(w).Encode(responses.ErrorResponse{Code: responses.BAD_TOKEN, Message: "User token not found or expired!"})
	}
	return isValidToken
}

func ParseBody(body io.ReadCloser, v interface{}) {
	reqBody, _ := ioutil.ReadAll(body)
	json.Unmarshal(reqBody, v)
}
