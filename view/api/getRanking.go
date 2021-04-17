package api

import "net/http"

func GetRanking(w http.ResponseWriter, req *http.Request) {
	JsonRequestResponder(w, `
{
	"Placeholder": "getRanking placeholder"
}
`, 200)
}