package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
)

func SearchUsers(w http.ResponseWriter, req *http.Request) {
	responder.JsonRequestResponder(w, `
{
	"Placeholder": "searchUsers placeholder"
}
`, 200)
}