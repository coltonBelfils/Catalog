package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func EditRanking(w http.ResponseWriter, req *http.Request) {
	responder.JsonRequestResponder(w, `
{
	"Placeholder": "editRanking placeholder"
}
`, 200)
}