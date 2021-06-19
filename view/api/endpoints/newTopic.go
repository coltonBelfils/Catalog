package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func NewTopic(w http.ResponseWriter, req *http.Request) {
	responder.JsonRequestResponder(w, `
{
	"Placeholder": "newTopic placeholder"
}
`, 200)
}