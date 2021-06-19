package middleware

import (
	"fmt"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			fmt.Println("options")
			OptionsResponder(w, 200)
			return
		}
		fmt.Println("serve")
		next.ServeHTTP(w, r)
	})
}