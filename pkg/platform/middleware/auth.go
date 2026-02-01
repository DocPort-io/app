package middleware

import (
	"log"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("running auth middleware")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
