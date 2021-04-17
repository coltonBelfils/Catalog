package api

import (
	"Catalog/niceErrors"
	"fmt"
	"net/http"
)

func JsonRequestResponder(w http.ResponseWriter, jsonData string, responseCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(responseCode)

	fmt.Fprint(w, jsonData)
}

func JsonRequestErrorResponder(w http.ResponseWriter, NiceError *niceErrors.NiceErrors, responseCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(responseCode)

	fmt.Fprint(w, NiceError.ToJson())
}

func HtmlRequestResponder(w http.ResponseWriter, htmlData string, responseCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "text/html")
	w.WriteHeader(responseCode)

	fmt.Fprint(w, htmlData)
}