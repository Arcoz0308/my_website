package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var global = map[string]*rateLimit{}

var globalLimit = 120 // number of request max for each minute

type rateLimit struct {
	Number int       // number of request already made
	Reset  time.Time // timestamp when the limit are reset
}

func CheckGlobalAPIRateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.RemoteAddr
		}
		if rl, ok := global[ip]; ok {
			if rl.Reset.Before(time.Now()) {
				log.Printf("nul")
				rl.Number = 1
				rl.Reset = time.Now().Add(time.Minute)
			} else if rl.Number < globalLimit {
				rl.Number = rl.Number + 1
			} else {
				// user are global rate limit
				log.Printf("address %s are global ratelimit for %f seconds", ip, time.Until(rl.Reset).Seconds())
				w.Header().Set("X-RateLimit-Global", "true")
				w.Header().Set("Retry-After", strconv.Itoa(int(time.Until(rl.Reset).Seconds())))
				w.Header().Set("X-RateLimit-Limit", strconv.Itoa(globalLimit))
				w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(globalLimit-rl.Number))
				w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%f", float64(rl.Reset.UnixMilli())/1000))
				w.Header().Set("X-RateLimit-Reset-After", fmt.Sprintf("%f", float64(time.Until(rl.Reset).Milliseconds())/1000))
				w.WriteHeader(http.StatusTooManyRequests)
				b, err := json.Marshal(struct {
					Message    string  `json:"message"`
					RetryAfter float64 `json:"retry_after"`
					Global     bool    `json:"global"`
				}{
					"You are being rate limited.",
					float64(time.Until(rl.Reset).Milliseconds()) / 1000,
					true,
				})
				if err != nil {
					io.WriteString(w, "internal server error")
					log.Panic(err)
				}
				io.WriteString(w, string(b))
				return
			}

		} else {
			global[ip] = &rateLimit{
				Number: 1,

				Reset: time.Now().Add(time.Minute),
			}
		}
		rl := global[ip]
		// set headers
		w.Header().Set("X-RateLimit-Limit", strconv.Itoa(globalLimit))
		w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(globalLimit-rl.Number))
		w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%f", float64(rl.Reset.UnixMilli())/1000))
		w.Header().Set("X-RateLimit-Reset-After", fmt.Sprintf("%f", float64(time.Until(rl.Reset).Milliseconds())/1000))

		next.ServeHTTP(w, r)
	})
}
