package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func GetHomepage(w http.ResponseWriter, req *http.Request) {
	responder.HtmlRequestResponder(w, req, "Homepage", "html/homepage.html")
}
