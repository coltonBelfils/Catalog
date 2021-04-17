package api

import "net/http"

func SearchUsers(w http.ResponseWriter, req *http.Request) {
	JsonRequestResponder(w, `
{
	"Placeholder": "searchUsers placeholder"
}
`, 200)
}