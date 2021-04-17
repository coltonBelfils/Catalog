package api

import "net/http"

func EditRanking(w http.ResponseWriter, req *http.Request) {
	JsonRequestResponder(w, `
{
	"Placeholder": "editRanking placeholder"
}
`, 200)
}