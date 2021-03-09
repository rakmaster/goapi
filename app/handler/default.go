package handler

import (
	"net/http"
)

// ShowDefault display default home page
func ShowDefault(res http.ResponseWriter, req *http.Request) {
	HTMLResponseWriter(res, http.StatusOK, "Hello World")
}
