package middleware

import (
	re "Catalog/view/api/responder"
	"fmt"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			fmt.Println("options")
			re.OptionsResponder(w, 200)
			return
		}
		fmt.Println("serve")
		next.ServeHTTP(w, r)
	})
}