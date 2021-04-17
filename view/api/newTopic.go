package api

import "net/http"

func NewTopic(w http.ResponseWriter, req *http.Request) {
	JsonRequestResponder(w, `
{
	"Placeholder": "newTopic placeholder"
}
`, 200)
}