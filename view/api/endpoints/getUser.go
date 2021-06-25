package endpoints

import (
	"Catalog/controler"
	"Catalog/niceErrors"
	"Catalog/view/api/responder"
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) { //the uuid of the user should be for internal uses only! get the user by username

	/* for later reference
	var jsonBody map[string]interface{}
	json.NewDecoder(r.Body).Decode(&jsonBody)
	*/

	if r.Method == "GET" {
		queryValues := r.URL.Query()
		username := queryValues.Get("username")

		gotUser, nGetErr := controler.GetUserByUsername(username)
		if nGetErr != nil {
			responder.JsonRequestErrorResponder(w, nGetErr, 400)
		}

		jsonConv, martialErr := json.Marshal(gotUser)
		if martialErr != nil {
			responder.JsonRequestErrorResponder(w, niceErrors.New(martialErr.Error(), "Unable to find user with username: "+username, niceErrors.JsonConvError, niceErrors.ERROR), 500)
		}

		responder.JsonRequestResponder(w, string(jsonConv), 200)

	} else {
		responder.JsonRequestErrorResponder(w, niceErrors.New("user called getUser with "+r.Method, r.Method+" calls are not allowed", niceErrors.InvalidActionByUserError, niceErrors.WARN), 405)
	}
}
