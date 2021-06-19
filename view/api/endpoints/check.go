package endpoints

import (
	"Catalog/view/api/responder"
	"net/http"
	"time"
)

func Check(w http.ResponseWriter, req *http.Request) {
	responder.JsonRequestResponder(w, `
{
	"check": "api is running",
	"start_time": "` + time.Now().String() + "\"\n}", 200)
}