package api

import (
	"net/http"
)

func Check(w http.ResponseWriter, req *http.Request) {
	JsonRequestResponder(w, `
{
	"check": "api is running"
}
`, 200)
}