package handler

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// ShowDefault display default home page
func ShowDefault(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	ResponseWriter(res, http.StatusOK, "Hello World", nil)
}
