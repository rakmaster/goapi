package controller

import (
	"net/http"

	"github.com/rakmaster/goapi/app/handler"
)

// ShowDefault display default home page
func ShowDefault(res http.ResponseWriter, req *http.Request) {
	handler.HTMLResponseWriter(res, http.StatusOK, "Hello World")
}
