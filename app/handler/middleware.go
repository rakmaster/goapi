package handler

import "net/http"

// JSONContentTypeMiddleware will add the json content type header
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/vnd.api+json; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}

// HTMLContentTypeMiddleware will add the html content type header
func HTMLContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}
