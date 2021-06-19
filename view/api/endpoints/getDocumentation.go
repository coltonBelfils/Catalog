package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func GetDocumentation(w http.ResponseWriter, req *http.Request) {
	responder.HtmlRequestResponder(w, req, "Documentation", "html/documentation.html")
}
