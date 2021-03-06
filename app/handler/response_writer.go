package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponseWriter will write result in http.ResponseWriter
func JSONResponseWriter(res http.ResponseWriter, statusCode int, data interface{}) error {
	res.WriteHeader(statusCode)
	httpResponse := NewResponse(data)
	err := json.NewEncoder(res).Encode(httpResponse)
	return err
}

// JAPIErrorResponseWriter writes a standard JSON:API error response
func JAPIErrorResponseWriter(res http.ResponseWriter, statusCode int, message string) error {
	res.WriteHeader(statusCode)
	errorResponse := NewErrorResponse(statusCode, message, message, message)
	err := json.NewEncoder(res).Encode(errorResponse)
	return err
}

// HTMLResponseWriter will write result in http.ResponseWriter
func HTMLResponseWriter(res http.ResponseWriter, statusCode int, message string) {
	res.WriteHeader(statusCode)
	fmt.Fprintf(res, "<html><head><meta charset='UTF-8'></head><body>%v</body></html>", message)
}
