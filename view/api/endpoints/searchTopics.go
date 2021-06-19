package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func SearchTopics(w http.ResponseWriter, req *http.Request) {
	responder.JsonRequestResponder(w, `
{
	"Placeholder": "searchTopics placeholder"
}
`, 200)
}