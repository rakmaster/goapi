package handler

import (
	"net/http"
)

// ShowDefault display default home page
func ShowDefault(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	ResponseWriter(res, http.StatusOK, "Hello World", nil)
}
