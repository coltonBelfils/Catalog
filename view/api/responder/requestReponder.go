package responder

import (
	"Catalog/niceErrors"
	"fmt"
	"net/http"
	"os"
	"time"
)

func JsonRequestResponder(w http.ResponseWriter, jsonData string, responseCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	w.Header().Set("content-type", "application/json")

	fmt.Fprint(w, jsonData)
}

func JsonRequestErrorResponder(w http.ResponseWriter, NiceError *niceErrors.NiceErrors, responseCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	w.Header().Set("content-type", "application/json")

	fmt.Fprint(w, NiceError.ToJson())
}

func HtmlRequestResponder(w http.ResponseWriter, req *http.Request, htmlTitle string, htmlFile string) *niceErrors.NiceErrors {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	w.Header().Set("content-type", "text/html")

	content, err := os.Open(htmlFile)
	if err != nil {
		return niceErrors.FromErrorFull(err, "couldn't open: " + htmlFile, "-", niceErrors.ERROR)
	}

	http.ServeContent(w, req, htmlTitle, time.Time{}, content)

	return nil
}

func OptionsResponder(w http.ResponseWriter, responseCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	w.WriteHeader(responseCode)
}