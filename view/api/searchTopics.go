package api

import "net/http"

func SearchTopics(w http.ResponseWriter, req *http.Request) {
	JsonRequestResponder(w, `
{
	"Placeholder": "searchTopics placeholder"
}
`, 200)
}