package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func EditTopic(w http.ResponseWriter, req *http.Request) {
	responder.JsonRequestResponder(w, `
{
	"Placeholder": "editTopic placeholder"
}
`, 200)
}