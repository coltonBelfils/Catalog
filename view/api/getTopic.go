package api

import "net/http"

func GetTopic(w http.ResponseWriter, req *http.Request) {
	JsonRequestResponder(w, `
{
	"Placeholder": "getTopic placeholder"
}
`, 200)
}