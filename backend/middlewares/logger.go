package middlewares

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

		crw := &customResponseWriter{ResponseWriter: w}
		// Call the next handler in the chain
		next.ServeHTTP(crw, r)

		// Log the response details
		fmt.Printf("Sent response %s\n", crw.body)
	})
}
