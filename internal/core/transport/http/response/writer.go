package response

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(rw http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: rw,
		statusCode:     -1,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *ResponseWriter) GetStatusCode() int {
	return rw.statusCode
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	if rw.statusCode == -1 {
		rw.statusCode = http.StatusOK
	}

	return rw.ResponseWriter.Write(b)
}
