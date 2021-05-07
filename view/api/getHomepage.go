package api

import (
	"net/http"
)

func GetHomepage(w http.ResponseWriter, req *http.Request) {
	HtmlRequestResponder(w, req, "Homepage", "html/homepage.html")
}
