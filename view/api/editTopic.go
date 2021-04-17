package api

import "net/http"

func EditTopic(w http.ResponseWriter, req *http.Request) {
	JsonRequestResponder(w, `
{
	"Placeholder": "editTopic placeholder"
}
`, 200)
}