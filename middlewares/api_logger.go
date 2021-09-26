package middlewares

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	StatusCode int
	Length     int
}

func (rw *statusWriter) Header() http.Header {
	return rw.ResponseWriter.Header()
}
func (rw *statusWriter) Write(p []byte) (int, error) {
	if rw.StatusCode == 0 {
		rw.StatusCode = 200
	}
	n, err := rw.ResponseWriter.Write(p)
	rw.Length += n
	return n, err

}
func (rw *statusWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
func LogAPIRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := statusWriter{ResponseWriter: w}
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(&sw, r)
		latency := time.Since(start).Milliseconds()
		ip := r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.RemoteAddr
		}
		// ip token [date] "method url" status (received bytes) (send bytes) latency
		msg := []string{
			ip,
			"-", // current not implement
			start.Format("[02/Jan/2006:15:04:05 Z0700]"),
			"\"" + r.Method,
			r.URL.RequestURI() + "\"",
			strconv.Itoa(sw.StatusCode),
			strconv.Itoa(int(r.ContentLength)) + "B",
			strconv.Itoa(int(sw.Length)) + "B",
			strconv.Itoa(int(latency)) + "ms",
		}
		log.Print(strings.Join(msg, " "))
	})
}
