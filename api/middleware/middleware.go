package middleware

import "net/http"

func HeaderMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(rw, r)
	})
}
