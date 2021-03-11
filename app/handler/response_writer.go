package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// JSONResponseWriter will write result in http.ResponseWriter
func JAPIResponseWriter(res http.ResponseWriter, statusCode int, data interface{}) error {
	res.WriteHeader(statusCode)
	httpResponse := NewJAPIResponse(data)
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

type Html struct {
	Title   string
	Message string
}

// HTMLResponseWriter will write result in http.ResponseWriter
func HTMLResponseWriter(res http.ResponseWriter, statusCode int, message string) {
	res.WriteHeader(statusCode)
	r := Html{
		Title:   "Simple Go JSON:API",
		Message: message,
	}
	parsedTemplate, _ := template.ParseFiles("html/index.html")
	err := parsedTemplate.Execute(res, r)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
