package middlewares

import "net/http"

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (w *customResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *customResponseWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...) // Capture the body
	return w.ResponseWriter.Write(b)
}
