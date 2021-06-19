package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func GetRanking(w http.ResponseWriter, req *http.Request) {
	responder.JsonRequestResponder(w, `
{
	"Placeholder": "getRanking placeholder"
}
`, 200)
}