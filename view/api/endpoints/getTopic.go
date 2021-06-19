package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func GetTopic(w http.ResponseWriter, req *http.Request) {
	responder.JsonRequestResponder(w, `
{
	"Placeholder": "getTopic placeholder"
}
`, 200)
}