package api

import "net/http"

func GetDocumentation(w http.ResponseWriter, req *http.Request) {
	HtmlRequestResponder(w, req, "Documentation", "html/documentation.html")
}
