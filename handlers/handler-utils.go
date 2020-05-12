package handlers

import "net/http"

// notFoundHandler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

}

// Log headers and body.
func HttpTraceAll(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do some trace things
		f.ServeHTTP(w, r)
	}
}
