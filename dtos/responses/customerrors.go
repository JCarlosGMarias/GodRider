package responses

import "fmt"

type ErrorCode int8

const (
	OK ErrorCode = iota
	METHOD_UNAUTHORIZED
	BAD_CREDENTIALS
	BAD_TOKEN
	REGISTER_NOT_FOUND
	ADD_ERROR
	WEBSERVICE_CONNECTION_FAILURE
)

func (response *ErrorResponse) Error() string {
	return fmt.Sprintf("%d: %s", response.Code, response.Message)
}
